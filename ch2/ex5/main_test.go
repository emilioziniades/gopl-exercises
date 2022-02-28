package main

import "testing"

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountLoop(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountLoop(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCount64Shift(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCount64Shift(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountRightClear(b *testing.B) {
	for i := 0; i < b.N; i++ {
		PopCountRightClear(0x1234567890ABCDEF)
	}
}
