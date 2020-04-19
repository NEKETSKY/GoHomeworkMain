package main

import (
	"fmt"
	"sort"
)

func main() {
	//case#1. Expected result: b, c, a
	input1 := map[int]string{2: "a", 0: "b", 1: "c"}
	fmt.Println(printSorted(input1))

	//case#2. Expected result: bb, aa, cc
	input2 := map[int]string{10: "aa", 0: "bb", 500: "cc"}
	fmt.Println(printSorted(input2))
}

//printSorted print map values sorted in order of increasing keys
func printSorted(a map[int]string) []string {
	key := make([]int, 0, len(a))
	var res []string
	for i := range a {
		key = append(key, i)
	}
	sort.Ints(key)
	for i := range key {
		res = append(res, a[key[i]])
	}
	return res
}