package main

import "fmt"

func main() {
	input := [...]int{1, 2, 3, 4, 5, 6}
	fmt.Println(average(input))
}

//average returns an average value of array
func average(a [6]int) float64 {
	var sum float64
	for i := range a {
		sum = sum + float64(a[i])
	}
	return sum/float64(len(a))
}
