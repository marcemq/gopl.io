// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 8.

// Echo4 prints its command-line arguments from position 0
package main

import (
	"fmt"
	"os"
	"strings"
)

//!+
func main() {
	fmt.Println(strings.Join(os.Args[:], " "))
}

//!-
