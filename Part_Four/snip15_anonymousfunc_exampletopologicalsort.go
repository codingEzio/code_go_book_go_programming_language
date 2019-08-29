package main

import (
	"fmt"
	"sort"
)

func main() {
	// prereqs(先行課) maps computer science courses to their prerequisites
	var prereqs = map[string][]string{
		"algorithms": {"data structures"},
		"calculus":   {"linear algebra"},

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

	for idx, course := range topoSort(prereqs) {
		fmt.Printf("%02d  %s\n", idx+1, course)
	}
}

func topoSort(mp map[string][]string) []string {
	/*
		Intro about topoSort
		- https://www.geeksforgeeks.org/depth-first-search-or-dfs-for-a-graph/
		- https://www.geeksforgeeks.org/topological-sorting/
	*/

	var order []string
	seen := make(map[string]bool)

	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true

				visitAll(mp[item])
				order = append(order, item)
			}
		}
	}

	var keys []string
	for key := range mp {
		keys = append(keys, key)
	}

	sort.Strings(keys)
	visitAll(keys)

	return order
}
