package main

import "fmt"

func main() {
	//case #1: different length. Expected result: three
	input1 := []string{"one", "two", "three"}
	fmt.Println(max(input1))

	//case #2: same length. Expected result: one
	input2 := []string{"one", "two"}
	fmt.Println(max(input2))
}

//max returns the longest word (first if there are more than one)
func max(a []string) string {
	var longest string
	for i := range a {
		if len(a[i]) > len(longest) {
			longest = a[i]
		}
	}
	return longest
}
