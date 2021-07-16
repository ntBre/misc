package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

const (
	help = `fctab is a tool for generating LaTeX tables from
fort.15,30,40 files. 
Flags:
`
)

func parseFlag() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(),
			"%s", help)
		flag.PrintDefaults()
	}
	flag.Parse()
}

var (
	rows    = flag.Int("r", 50, "maximum number of rows in a table")
	cols    = flag.Int("c", 3, "maximum number of columns in a table")
	caption = flag.String("caption", "", "caption template for the tables")
)

// Read15 reads a fort.15 file and returns a slice of the force
// constants formatted for a LaTeX table
func Read15(filename string) (fcs []string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var (
		head bool = true
		dim  int
	)
	row, col := 1, 1
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if head {
			d, _ := strconv.Atoi(fields[0])
			dim = d * 3
			head = false
			continue
		}
		for _, v := range fields {
			f, _ := strconv.ParseFloat(v, 64)
			fcs = append(fcs,
				fmt.Sprintf("$F_{%d,%d}$ & %.6f",
					row, col, f))
			if col <= dim {
				col++
			} else {
				col = 1
				row++
			}
		}
	}
	return
}

// DimMap3 maps positions in a 1D array to the corresponding cubic
// force constant subscript
func DimMap3(dim int) map[int][]int {
	conv := make(map[int][]int, 0)
	for i := 1; i <= dim; i++ {
		for j := 1; j <= i; j++ {
			for k := 1; k <= j; k++ {
				list := []int{i, j, k}
				sort.Ints(list)
				a, b, c := list[0], list[1], list[2]
				key := a + (b-1)*b/2 + (c-1)*c*(c+1)/6
				conv[key] = list
			}
		}
	}
	return conv
}

// Read30 reads a fort.30 file and returns a slice of the force
// constants formatted for a LaTeX table
func Read30(filename string) (fcs []string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var (
		head bool = true
		dim  int
		conv map[int][]int
	)
	count := 1
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if head {
			d, _ := strconv.Atoi(fields[0])
			dim = 3 * d
			conv = DimMap3(dim)
			head = false
			continue
		}
		for _, v := range fields {
			list := conv[count]
			f, _ := strconv.ParseFloat(v, 64)
			fcs = append(fcs,
				fmt.Sprintf("$F_{%d,%d,%d}$ & %.4f",
					list[0], list[1], list[2], f))
			count++
		}
	}
	return
}

// DimMap4 maps positions in a 1D array to the corresponding quartic
// force constant subscript
func DimMap4(dim int) map[int][]int {
	conv := make(map[int][]int, 0)
	for i := 1; i <= dim; i++ {
		for j := 1; j <= i; j++ {
			for k := 1; k <= j; k++ {
				for l := 1; l <= k; l++ {
					list := []int{i, j, k, l}
					sort.Ints(list)
					a, b, c, d := list[0], list[1], list[2], list[3]
					key := a + (b-1)*b/2 + (c-1)*c*(c+1)/6 +
						(d-1)*d*(d+1)*(d+2)/24
					conv[key] = list
				}
			}
		}
	}
	return conv
}

// Read40 reads a fort.40 file and returns a slice of the force
// constants formatted for a LaTeX table
func Read40(filename string) (fcs []string) {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(f)
	var (
		head bool = true
		dim  int
		conv map[int][]int
	)
	count := 1
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		if head {
			if len(fields) < 2 {
				fmt.Fprintf(os.Stderr, "Header format error in %s\n",
					filename)
				os.Exit(1)
			}
			d, _ := strconv.Atoi(fields[0])
			dim = 3 * d
			conv = DimMap4(dim)
			head = false
			continue
		}
		for _, v := range fields {
			list := conv[count]
			f, _ := strconv.ParseFloat(v, 64)
			fcs = append(fcs,
				fmt.Sprintf("$F_{%d,%d,%d,%d}$ & %.2f",
					list[0], list[1], list[2], list[3], f))
			count++
		}
	}
	return
}

func header(cont bool) string {
	var buf bytes.Buffer
	buf.WriteString(`\begin{table}[ht]
\centering
`)
	fmt.Fprintf(&buf, "\\caption{%s", *caption)
	if cont {
		fmt.Fprintln(&buf, " (cont.)}")
	} else {
		fmt.Fprintln(&buf, "}")
	}
	buf.WriteString(`\begin{tabular}{|`)
	for i := 0; i < *cols; i++ {
		fmt.Fprint(&buf, "cc|")
	}
	fmt.Fprint(&buf, "}\n")
	return buf.String()
}

const footer = `\end{tabular}
\end{table}
`

func main() {
	parseFlag()
	dir := "."
	args := flag.Args()
	if len(args) == 1 {
		dir = args[0]
	}
	files := []string{"fort.15", "fort.30", "fort.40"}
	funcs := []func(string) []string{Read15, Read30, Read40}
	fcs := make([]string, 0)
	for i, file := range files {
		file = filepath.Join(dir, file)
		if _, err := os.Stat(file); os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "file %s does not exist, aborting\n",
				file)
			os.Exit(1)
		}
		fcs = append(fcs, funcs[i](file)...)
	}
	fmt.Print(header(false))
	count := 1
	var buf bytes.Buffer
	var chunk []string
	for i := 0; i+*cols <= len(fcs) || (len(fcs)-i < *cols && len(fcs)-i > 0); i += *cols {
		if count == *rows {
			fmt.Println(footer)
			fmt.Print(header(true))
			count = 1
		}
		if len(fcs)-i < *cols {
			chunk = fcs[i:]
		} else {
			chunk = fcs[i : i+*cols]
		}
		for f := range chunk {
			fmt.Fprint(&buf, chunk[f])
			if f == *cols-1 {
				fmt.Fprint(&buf, "\\\\\n")
			} else if f == len(chunk)-1 && f < *cols-1 {
				for ; f < *cols-1; f++ {
					fmt.Fprint(&buf, " & ")
				}
				fmt.Fprint(&buf, "\\\\\n")
			} else {
				fmt.Fprint(&buf, " & ")
			}
		}
		count++
		fmt.Print(buf.String())
		buf.Reset()
	}
	fmt.Println(footer)
}
