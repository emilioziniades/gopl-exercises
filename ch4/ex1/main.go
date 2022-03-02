package main

import (
	"crypto/sha256"
	"fmt"
	"math/bits"
)

func main() {
	fmt.Println(BitsDiffSHA("x", "X"))
}

func BitsDiffSHA(s1, s2 string) int {
	count := 0
	c1 := sha256.Sum256([]byte(s1))
	c2 := sha256.Sum256([]byte(s2))

	for i := 0; i < len(c1); i++ {
		i1 := c1[i]
		i2 := c2[i]
		diff := i1 ^ i2
		count += bits.OnesCount8(diff)
	}

	return count
}
