package main

import (
	"fmt"
	"math"
)

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc[byte(x>>(0*8))] +
		pc[byte(x>>(1*8))] +
		pc[byte(x>>(2*8))] +
		pc[byte(x>>(3*8))] +
		pc[byte(x>>(4*8))] +
		pc[byte(x>>(5*8))] +
		pc[byte(x>>(6*8))] +
		pc[byte(x>>(7*8))])
}

func PopCountLoop(x uint64) (res int) {
	for i := 0; i < 8; i++ {
		res += int(pc[byte(x>>(i*8))])
	}
	return
}

func PopCount64Shift(x uint64) (res int) {
	for i := 0; i < 64; i++ {
		res += int(x >> i & 1)
	}
	return
}

func PopCountRightClear(x uint64) (res int) {
	for x != 0 {
		res++
		x = x & (x - 1)
	}
	return
}

func main() {
	val := uint64(math.MaxUint64)
	fmt.Println(PopCount(val))
	fmt.Println(PopCountLoop(val))
	fmt.Println(PopCount64Shift(val))
	fmt.Println(PopCountRightClear(val))
}
