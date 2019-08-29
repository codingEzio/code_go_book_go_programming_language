package main

import (
	"fmt"
	"os"
	"strings"
)

// prereqs maps computer science courses to their prerequisites.
var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	// new course (& a cycle)
	"linear algebra": {"calculus"},

	// another cycle
	"intro to programming": {"data structures"},

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
	order, err := topoSort2(prereqs)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	for idx, course := range order {
		fmt.Printf("%02d: %s\n", idx+1, course)
	}
}

func topoSort2(mp map[string][]string) (order []string, err error) {
	resolved := make(map[string]bool)

	var visitAll func([]string, []string)
	visitAll = func(items []string, parents []string) {
		for _, item := range items {
			itemResolved, seen := resolved[item]

			if seen && !itemResolved {
				// ignore error since `item` has be in parents
				start, _ := index(item, parents)

				err = fmt.Errorf("cycle: %s",
					strings.Join(append(parents[start:], item), " -> "))
			}

			if !seen {
				resolved[item] = false
				visitAll(mp[item], append(parents, item))
				resolved[item] = true

				order = append(order, item)
			}
		}
	}

	for key := range mp {
		if err != nil {
			return nil, err
		}
		visitAll([]string{key}, nil)
	}

	return order, nil
}

func index(str string, slice []string) (int, error) {
	for idx, item := range slice {
		if str == item {
			return idx, nil
		}
	}
	return 0, fmt.Errorf("not found")
}
