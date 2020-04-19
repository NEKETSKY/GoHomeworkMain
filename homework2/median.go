package main

import (
	"fmt"
	"sort"
)

func median(i []int) float64 {
	sort.Ints(i)
	if len(i)%2 == 0 {
		return float64(i[len(i)/2]+i[len(i)/2-1]) / 2
	}

	return float64(i[(len(i) / 2)])
}

func main() {
	// Case #1: even amount of elements. Expected value: 5.5
	test := []int{5, 3, 6, 8, 4, 9}
	fmt.Println("First example:", median(test))

	// Case #2: odd amount of elements. Expected value: 7
	test = []int{4, 3, 7, 9, 15}
	fmt.Println("Second example:", median(test))
}
