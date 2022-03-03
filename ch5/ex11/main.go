package main

import (
	"fmt"
	"log"
	"sort"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},
	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	sortedList, err := topoSort(prereqs)
	if err != nil {
		log.Fatalln(err)
	}
	for i, course := range sortedList {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) ([]string, error) {

	var cyclic bool
	var order []string
	seen := make(map[string]bool)
	recStack := make(map[string]bool)
	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if seen[item] && recStack[item] {
				cyclic = true
			}

			if !recStack[item] {
				recStack[item] = true
			}

			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}

			recStack[item] = false
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	if cyclic {
		return nil, fmt.Errorf("Cycle detected")
	}
	return order, nil
}
