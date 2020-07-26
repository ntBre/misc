package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintln(os.Stderr, "no expression present")
		os.Exit(1)
	}
	re, err := regexp.Compile(os.Args[1])
	if err != nil {
		fmt.Fprintln(os.Stderr, "error parsing expression")
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(re.MatchString(line))
	}
}
