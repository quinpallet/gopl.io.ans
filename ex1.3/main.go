// © 2019 kenqlo
// GOPL: 練習問題 1.3

// ページ9

// Echo2 and Echo3
package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var diff int64
	var comparison string

	// Echo2
	start := time.Now()
	fmt.Println("echo2: " + echo2())
	elapsed2 := time.Since(start).Nanoseconds()
	fmt.Printf("%d nsec elapsed\n", elapsed2)

	// Echo3
	start = time.Now()
	fmt.Println("echo3: " + echo3())
	elapsed3 := time.Since(start).Nanoseconds()
	// differences
	diff = elapsed2 - elapsed3
	if diff > 0 {
		comparison = "(" + strconv.FormatInt(diff, 10) + " nsec faster)"
	} else if diff > 0 {
		comparison = "(" + strconv.FormatInt(diff, 10) + " nsec slower)"
	}
	fmt.Printf("%d nsec elapsed %s\n", elapsed3, comparison)
}

func echo2() string {
	// Echo2 version
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	// Echo3 version
	return strings.Join(os.Args[1:], " ")
}

//!-
