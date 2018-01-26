package main

import (
	"fmt"
	"sort"
)

var prereqs = map[string][]string{
	"linear algebra": {"calculus"},
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
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
	ts := topoSort(prereqs)
	for i, course := range ts {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
	fmt.Printf("Topological Orderings: %v\n", isTopologicalOrdered(ts))
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func(items []string)
	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range m {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)
	return order
}

func isTopologicalOrdered(ts []string) error {
	nodes := make(map[string]int)

	for i, course := range ts {
		nodes[course] = i
	}
	for course, i := range nodes {
		nodes[course] = i
	}

	for course, i := range nodes {
		for _, prereq := range prereqs[course] {
			if i < nodes[prereq] {
				return fmt.Errorf("%s,%s are cycled\n", course, prereq)
			}
		}
	}
	return nil
}
