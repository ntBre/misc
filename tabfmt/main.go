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
	text := strings.Split(
		strings.TrimSpace(input), "\n")
	var split []string
	for _, line := range text {
		line = strings.ReplaceAll(line, "\\&", amp)
		split = strings.Split(line, "&")
		switch {
		case strings.Contains(line, "multicolumn"):
			for c := range split {
				split[c] = strings.TrimSpace(split[c])
			}
		case strings.Contains(line, "\\begin{tabular}"):
		case strings.Contains(line, "\\end{tabular}"):
		default:
			// extend widths by the difference
			ls := len(split)
			lw := len(widths)
			if ls > lw {
				widths = append(widths, make([]int, ls-lw)...)
			}
			for c := range split {
				split[c] = strings.TrimSpace(split[c])
				split[c] = strings.ReplaceAll(split[c], amp, "\\&")
				if clen := len(split[c]); clen > widths[c] {
					widths[c] = clen
				}
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
		if len(line) == 1 && strings.Contains(line[0], "\\hline") {
			fmt.Fprintln(w, line[0])
			continue
		}
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
	WriteTab(os.Stdout, lines, widths)
}
