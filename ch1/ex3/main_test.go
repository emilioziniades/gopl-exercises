package main

import "testing"

var testSlice = []string{
	"this",
	"is",
	"a",
	"slice",
	"of",
	"strings",
	"to",
	"test",
	"the",
	"echo",
	"functions",
	"but",
	"it",
	"seems",
	"that",
	"I",
	"should",
	"add",
	"more",
	"lines",
	"to",
	"improve",
	"the",
	"benchmark",
	"accuracy",
}

func BenchmarkEchoNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoNaive(testSlice)
	}
}

func BenchmarkEchoBetter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EchoNaive(testSlice)
	}
}
