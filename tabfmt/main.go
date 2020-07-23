package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

const amp = "amp;"

// ParseTab parses an input string into its component lines and
// returns those lines as a [][]string, along with the maximum widths
// of each column as a []int
func ParseTab(input string) (lines [][]string, widths []int) {
	text := strings.Split(input, "\n")
	start := true
	for _, line := range text {
		line = strings.ReplaceAll(line, "\\&", amp)
		split := strings.Split(line, "&")
		if strings.Contains(line, "multicolumn") {
			for c := range split {
				split[c] = strings.TrimSpace(split[c])
			}
			lines = append(lines, split)
			continue
		}
		if start {
			widths = make([]int, len(split), len(split))
			start = false
		}
		for c := range split {
			split[c] = strings.TrimSpace(split[c])
			split[c] = strings.ReplaceAll(split[c], amp, "\\&")
			if clen := len(split[c]); clen > widths[c] {
				widths[c] = clen
			}
		}
		lines = append(lines, split)
	}
	return
}

// WriteTab writes a formatted table to w
func WriteTab(w io.Writer, lines [][]string, widths []int) {
	var buf bytes.Buffer
	for _, line := range lines {
		buf.Reset()
		for c, col := range line {
			w := strconv.Itoa(widths[c])
			if c == 0 {
				fmt.Fprintf(&buf, "%-"+w+"s", col)
			} else {
				fmt.Fprintf(&buf, "%"+w+"s", col)
			}
			if c < len(line)-1 {
				fmt.Fprint(&buf, " & ")
			}
		}
		buf.WriteString("\n")
		fmt.Fprint(w, buf.String())
	}
}

func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	lines, widths := ParseTab(string(input))
	WriteTab(os.Stdout, lines[:len(lines)-1], widths)
}
