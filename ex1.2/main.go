// © 2019 kenqlo
// GOPL: 練習問題 1.2

// ページ9

// Echo2 based prints its command-line arguments with the indexes.
package main

import (
	"fmt"
	"os"
)

func main() {
	for idx, arg := range os.Args[1:] {
		fmt.Println(idx, arg)
	}
}

//!-
