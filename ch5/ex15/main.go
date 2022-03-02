package main

import "fmt"

func main() {
	nums := []int{-12, -35, 22, 21}
	fmt.Printf("Min of %v: %v\n", nums, min(nums...))
	fmt.Printf("Max of %v: %v\n", nums, max(nums...))

	fmt.Println(min())
	fmt.Println(max())

	fmt.Println(min_nz(3, -1, 4))
	fmt.Println(max_nz(3, -1, 4))

	fmt.Println(min_nz())
	fmt.Println(max_nz())

}

func min(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	min := vals[0]
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max(vals ...int) int {
	if len(vals) == 0 {
		return 0
	}
	max := vals[0]
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}

func min_nz(first int, vals ...int) int {
	min := first
	for _, val := range vals {
		if val < min {
			min = val
		}
	}
	return min
}

func max_nz(first int, vals ...int) int {
	max := first
	for _, val := range vals {
		if val > max {
			max = val
		}
	}
	return max
}
