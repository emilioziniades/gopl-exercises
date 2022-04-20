package main

import (
	"fmt"
	"testing"
)

func TestIntSet(t *testing.T) {
	var x IntSet
	var y = make(Set)

	x.Add(91)
	y.Add(91)

	if x.Has(91) != y.Has(91) {
		t.Errorf("IntSet.Has(1) != Set.Has(1)")
	}
	if x.Has(87) != y.Has(87) {
		t.Errorf("IntSet.Has(91) != Set.Has(91)")
	}
}

func TestIntSetUnion(t *testing.T) {
	var s1 = make(Set)
	var s2 = make(Set)
	var is1, is2 IntSet

	nums := []int{5, 7}
	unionNums := []int{11, 13}

	for _, e := range nums {
		s1.Add(e)
		is1.Add(e)
	}

	for _, e := range unionNums {
		s2.Add(e)
		is2.Add(e)
	}

	s1.UnionWith(s2)
	is1.UnionWith(&is2)

	for _, e := range append(nums, unionNums...) {
		if !(s1.Has(e) && is1.Has(e)) {
			t.Errorf("Set.Has(%v)  or IntSet.Has(%v) = false", e, e)
		}
	}
}

func Example_one() {
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
