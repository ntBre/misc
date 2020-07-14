// TODO try extracting more information from the fields
// - Subject field in particular seems to have a lot of info
//  - but unclear if it's regular enough to parse
package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

const (
	progname = "pdfinfgo"
)

type ref struct {
	name    string
	author  string
	title   string
	journal string
	volume  string
	pages   string
	year    string
}

func (r ref) String() string {
	var buf bytes.Buffer
	fmt.Fprintf(&buf, "@article{%s,\n", r.name)
	fmt.Fprintf(&buf, "Author={%s},\n", r.author)
	fmt.Fprintf(&buf, "Title={%s},\n", r.title)
	fmt.Fprintf(&buf, "Journal={%s},\n", r.journal)
	fmt.Fprintf(&buf, "Volume={%s},\n", r.volume)
	fmt.Fprintf(&buf, "Pages={%s},\n", r.pages)
	fmt.Fprintf(&buf, "Year={%s}}", r.year)
	return buf.String()
}

func pdfinfo(pdfname string) string {
	cmd := exec.Command("pdfinfo", pdfname)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v reading from pdfinfo\n", progname, err)
		os.Exit(1)
	}
	scanner := bufio.NewScanner(stdout)
	err = cmd.Start()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v starting pdfinfo\n", progname, err)
		os.Exit(1)
	}
	var (
		title   = "Title:"
		subject = "Subject:"
		author  = "Author:"
	)
	r := new(ref)
	refname := path.Base(pdfname)
	r.name = refname[:len(refname)-len(path.Ext(refname))]
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case strings.Contains(line, title):
			r.title = strings.TrimSpace(line[len(title):])
		case strings.Contains(line, subject):
			r.journal = strings.TrimSpace(line[len(subject):])
		case strings.Contains(line, author):
			r.author = strings.TrimSpace(line[len(author):])
		}
	}
	err = cmd.Wait()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: %v running pdfinfo\n", progname, err)
		os.Exit(1)
	}
	return r.String()
}

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: too few arguments\n", progname)
		os.Exit(1)
	}
	pdfname := os.Args[1]
	if _, err := os.Stat(pdfname); os.IsNotExist(err) {
		fmt.Fprintf(os.Stderr, "%s: input file %q does not exist\n", progname, pdfname)
		os.Exit(1)
	}
	refstring := pdfinfo(pdfname)
	fmt.Println(refstring)
}
