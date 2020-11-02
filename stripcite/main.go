package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "%s: not enough input arguments\n", os.Args[0])
		os.Exit(2)
	}
	infile, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s: file %q does not exist\n", os.Args[0], infile)
		os.Exit(1)
	}
	var (
		incite bool
		sid    int
		eid    int
	)
	scanner := bufio.NewScanner(infile)
	start := "\\cite{"
	end := "}"
	for scanner.Scan() {
		line := scanner.Text()
		// either found a new citation or already in citation
		if sid = strings.Index(line, start); incite || sid > 0 {
			// found the end of a citation
			if eid = strings.Index(line, end); eid > 0 {
				if sid < 0 {
					sid = 0
				}
				fmt.Println(line[sid:eid+1])
				incite = false
			} else {
				fmt.Print(line[sid:])
				incite = true
			}
		}
	}
}
