// TODO handle lines with nothing in the first column
// - TrimLeft(line, " |") trims not only "  |" but also "  |     |" in these cases
// - could compare to colfmt line or something produced in that step
package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"regexp"
	"strings"
)

const header = `\begin{table}[ht]
\centering
\caption{}
\label{tab:}
\begin{tabular}`

const footer = `\end{tabular}
\end{table}
`

var (
	print = flag.Bool("p", false, "instead of writing the files print to stdout")
)

func colfmt(hline string) (int, string) {
	var (
		buf    bytes.Buffer
		newcol bool
		ncol   int
	)
	buf.WriteString("{")
	for _, c := range hline {
		if c == '|' || c == '+' {
			buf.WriteString("|")
			newcol = true
		} else if c == '-' && newcol {
			buf.WriteString("c")
			ncol++
			newcol = false
		}
	}
	buf.WriteString("}")
	return ncol, buf.String()
}

func main() {
	var (
		line    string
		dir     string
		outfile string
		col     string
		ncol    int
		tbl     bool
		next    bool
		buf     bytes.Buffer
		scanner *bufio.Scanner
	)
	flag.Parse()
	args := flag.Args()
	if len(args) == 1 {
		filename := args[0]
		dir = path.Dir(filename)
		f, err := os.Open(filename)
		if err != nil {
			fmt.Fprintf(os.Stderr, "orgtab: %v\n", err)
			os.Exit(1)
		}
		scanner = bufio.NewScanner(f)
	} else {
		dir = "."
		// https://stackoverflow.com/questions/22744443/check-if-there-is-something-to-read-on-stdin-in-golang
		if stat, _ := os.Stdin.Stat(); stat.Mode()&os.ModeCharDevice != 0 {
			fmt.Fprintf(os.Stderr, "orgtab: empty stdin\n")
			os.Exit(1)
		}
		scanner = bufio.NewScanner(os.Stdin)
	}
	empty := regexp.MustCompile(`^\s*$`)
	for scanner.Scan() {
		line = scanner.Text()
		switch {
		case strings.Contains(line, "#+tblname:"):
			next = true
			fields := strings.Fields(line)
			if len(fields) < 2 {
				panic("tblname not found")
			}
			outfile = dir + "/" + fields[1] + ".tex"
		case next:
			ncol, col = colfmt(line)
			fmt.Fprintln(&buf, header+col)
			fmt.Fprintln(&buf, "\\hline")
			next = false
			tbl = true
		case tbl && empty.MatchString(line):
			tbl = false
			fmt.Fprintln(&buf, footer)
			if !*print {
				err := ioutil.WriteFile(outfile, buf.Bytes(), 0755)
				if err != nil {
					panic(err)
				}
			} else {
				fmt.Print(buf.String())
			}
			buf.Reset()
		case tbl && strings.Contains(line, "---"):
			fmt.Fprintln(&buf, "\\hline")
		case tbl:
			line = strings.TrimLeft(line, " |")
			line = strings.TrimRight(line, " |")
			line = strings.ReplaceAll(line, "|", "&")
			if split := strings.Split(line, "&"); len(split) < ncol {
				line = " & " + line
			}
			fmt.Fprintln(&buf, line+"\\\\")
		}
	}
}
