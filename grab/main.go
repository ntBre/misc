// (progn (setq compile-command "go build . && scp grab woods:.") (recompile))
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

var (
	base = "/tmp/jax/forbrent/pts/inp"
)

// process a file, extracting the geometry and energy and storing the
// result in res
func Process(file string) (result Res, err error) {
	f, err := os.Open(file)
	defer f.Close()
	if err != nil {
		return
	}
	scanner := bufio.NewScanner(f)
	var (
		line string
		geom bool
		str  strings.Builder
	)
	space := regexp.MustCompile(` +`)
	for scanner.Scan() {
		line = scanner.Text()
		switch {
		case strings.Contains(line, "geometry="):
			geom = true
		case geom && strings.Contains(line, "}"):
			result.Geom = str.String()
			geom = false
		case geom:
			if len(strings.Fields(line)) < 4 {
				continue
			}
			line = space.ReplaceAllString(line, " ")
			fmt.Fprintln(&str, strings.TrimLeft(line, " "))
		case strings.Contains(line, "energy= "):
			fields := strings.Fields(line)
			result.Val, err = strconv.ParseFloat(fields[2], 64)
			if err != nil {
				return
			}
		}
	}
	return
}

type Res struct {
	Geom string
	Val  float64
}

func main() {
	if len(os.Args) == 2 {
		base = os.Args[1]
	}
	files, _ := filepath.Glob(
		filepath.Join(base, "job.*.out"),
	)
	files = append(files, filepath.Join(base, "ref.out"))
	hold := make([]Res, len(files))
	var wg sync.WaitGroup
	ch := make(chan struct{}, 40)
	for i, f := range files {
		wg.Add(1)
		ch <- struct{}{}
		go func(i int, f string) {
			defer func() {
				<-ch
				wg.Done()
			}()
			r, err := Process(f)
			if err != nil {
				fmt.Println(err)
				return
			}
			hold[i] = r
		}(i, f)
		wg.Wait()
	}
	res := make(map[string]float64)
	for _, r := range hold {
		if r != (Res{}) {
			res[r.Geom] = r.Val
		}
	}
	byt, _ := json.MarshalIndent(res, "", "\t")
	fmt.Println(string(byt))
}
