package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	incrementer := 0
	gs := 100

	for i := 0; i < gs; i++ {
		wg.Add(1)
		go func() {
			incrementer++
			fmt.Println(incrementer)
			wg.Done()
		}()
		wg.Wait()
	}

	fmt.Println("end value:", incrementer)
}
