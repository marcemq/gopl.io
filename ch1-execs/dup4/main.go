// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 13.
//!+

// Dup4 prints the name of all files in wich duplicated lines occurs.
// It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			if countLines(f) {
				fmt.Println(arg)
			}
			f.Close()
		}
	}
}

func countLines(f *os.File) bool {
	counts := make(map[string]int)
	flag := false
	input := bufio.NewScanner(f)
	for input.Scan() {
		counts[input.Text()]++
		if counts[input.Text()] > 1 {
			flag = true
			break
		}
	}
	return flag
	// NOTE: ignoring potential errors from input.Err()
}

//!-
