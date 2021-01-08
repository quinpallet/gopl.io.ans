// © 2019 kenqlo
// GOPL: 練習問題 1.1

// ページ9

// Echo3 based prints its command-line arguments with the command name.
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
