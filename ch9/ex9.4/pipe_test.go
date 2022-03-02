package main

import (
	"fmt"
	"testing"
	"time"
)

var tests = make([]int, 31)

func init() {
	for i := range tests {
		tests[i] = 1 << i
	}
}

func TestMakePipeline(t *testing.T) {
	for _, n := range tests {
		testMakePipelineN(n, t)
	}
}

func testMakePipelineN(n int, t *testing.T) {
	start := time.Now()
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("got to recovered panic at routine", n)
			t.Fatalf("%d routines failed: %v\n", n, r)
		}
		fmt.Printf("%d routines passed in %s\n", n, time.Since(start))
	}()
	makepipeline(n)
}
