package main

import "fmt"

func main() {
	input := []int64{1, 2, 5, 15}
	fmt.Println(reverse(input))
}

//reverse returns the copy of the slice in reverse order
func reverse(a []int64) []int64 {
	var rvrs []int64
	for i := range a {
		rvrs = append(rvrs, a[len(a)-1-i])
	}
	return rvrs
}
