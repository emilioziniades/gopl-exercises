package main

import "testing"

func BenchmarkMandelbrot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Generate(false)
	}
}

func BenchmarkMandelbrotParallell(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateParallell(false)
	}
}
