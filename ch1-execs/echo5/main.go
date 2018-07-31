// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.
//!+

// Echo6 prints its command-line index and arguments
package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	s, bsep, psep := "", " ", ":"
	for ind, arg := range os.Args[1:] {
		s += strconv.Itoa(ind+1) + psep + arg + bsep
	}
	fmt.Println(s)
}

//!-
