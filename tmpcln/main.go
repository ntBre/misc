package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"sort"
	"strconv"
	"strings"
	"sync"
)

const (
	nodeMin = 1
	nodeMax = 116
	nodeFmt = "%03d"
	nodePfx = "cn"
)

var (
	count = flag.Bool("c", false, "Count tmp usage but do not remove")
	dir   = flag.String("d", "/tmp/", "Specify the tmp directory to clean")
	user  = flag.String("u", "$USER", "Specify the user")
	dry   = flag.Bool("y", false, "Do a dry run, printing the commands to be run but not executing them")
)

func clean(node string, wg *sync.WaitGroup) (bytes int) {
	defer wg.Done()
	du := exec.Command("ssh", node, "-t", "du", "-bs", *dir+*user)
	rm := exec.Command("ssh", node, "-t", "rm", "-rf", *dir+*user)
	if *dry {
		fmt.Println(du.String())
		fmt.Println(rm.String())
		return
	}
	out, err := du.Output()
	if err != nil {
		// might want to panic if error other than "No such file or directory"
		return
	}
	// Example output of "du -bs School/": 7311459760      School/
	bytes, _ = strconv.Atoi(strings.Fields(string(out))[0])
	if *count {
		return
	}
	if bytes > 0 {
		err := rm.Run()
		if err != nil {
			panic(err)
		}
	}
	return
}

func prompt() {
	if !(*dry || *count) {
		fmt.Print("Are you sure you want to delete all tmp directories?\n")
		fmt.Print("This will disrupt all running jobs. [y/N] ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		switch scanner.Text() {
		case "Y", "y", "yes":
			return
		default:
			fmt.Println("Aborting")
			os.Exit(1)
		}
	}
}

type LockMap struct {
	sync.Mutex
	Map map[string]int
}

func NewLockMap() *LockMap {
	l := new(LockMap)
	l.Map = make(map[string]int)
	return l
}

func main() {
	flag.Parse()
	var (
		node  string
		wg    sync.WaitGroup
		bytes int
	)
	prompt()
	nodeMap := NewLockMap()
	*dry = true
	for i := nodeMin; i <= nodeMax; i++ {
		node = fmt.Sprintf(nodePfx+nodeFmt, i)
		wg.Add(1)
		go func(node string) {
			b := clean(node, &wg)
			if b > 0 {
				nodeMap.Lock()
				nodeMap.Map[node] = b
				nodeMap.Unlock()
				bytes += b
			}
		}(node)
	}
	wg.Wait()
	temp := make([]string, 0, len(nodeMap.Map))
	for k := range nodeMap.Map {
		temp = append(temp, k)
	}
	sort.Strings(temp)
	for _, node := range temp {
		fmt.Printf("%s: %d bytes\n", node, nodeMap.Map[node])
	}
}
