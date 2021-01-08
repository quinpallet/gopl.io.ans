// © 2019 kenqlo
// GOPL: 練習問題 1.3

// ページ9

// Echo2 and Echo3
package main

import (
	"testing"
)

func BenchmarkEcho2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo2()
	}
}

func BenchmarkEcho3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		echo3()
	}
}

//!-
