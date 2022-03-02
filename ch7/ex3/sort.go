package main

import (
	"bytes"
	"fmt"
)

type tree struct {
	value       int
	left, right *tree
}

func Sort(values []int) {
	var root *tree
	for _, v := range values {
		root = add(root, v)
	}
	appendValues(values[:0], root)
}

func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}
	return values
}

func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}
	if value < t.value {
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}
	return t
}

func (t *tree) String() string {
	order := make([]int, 0)
	order = appendValues(order, t)
	if len(order) == 0 {
		return "[]"
	}
	buf := new(bytes.Buffer)
	fmt.Fprintf(buf, "[%d", order[0])
	for _, o := range order[1:] {
		fmt.Fprintf(buf, " %d", o)
	}
	fmt.Fprintf(buf, "]")
	return buf.String()

}
func main() {
	root := &tree{value: 3}
	root = add(root, 2)
	root = add(root, 4)
	fmt.Println(root)
}
