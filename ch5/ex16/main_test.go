package main

import (
	"fmt"
	"testing"
)

var a10 [10]string
var a100 [100]string

func init() {
	for i := 0; i < len(a10); i++ {
		a10[i] = "a"
	}

	for i := 0; i < len(a100); i++ {
		a100[i] = "a"
	}
	fmt.Println(a10, a100)
}

func benchmarkJoin(a []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Join(",", a...)
	}
}

func BenchmarkJoin10(b *testing.B)  { benchmarkJoin(a10[:], b) }
func BenchmarkJoin100(b *testing.B) { benchmarkJoin(a100[:], b) }

func benchmarkJoinTorbiak(a []string, b *testing.B) {
	for n := 0; n < b.N; n++ {
		JoinTorbiak(",", a...)
	}
}

func BenchmarkJoinTorbiak10(b *testing.B)  { benchmarkJoinTorbiak(a10[:], b) }
func BenchmarkJoinTorbiak100(b *testing.B) { benchmarkJoinTorbiak(a100[:], b) }
