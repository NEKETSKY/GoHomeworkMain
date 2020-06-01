package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	start := time.Now()
	var t *time.Timer
	initreset := make(chan bool)
	t = time.AfterFunc(randomDuration(), func() {
		fmt.Println(time.Since(start))
		initreset <- true
	})
	for time.Since(start) < 5*time.Second {
		<-initreset
		t.Reset(randomDuration())
	}
}

func randomDuration() time.Duration {
	return time.Duration(rand.Int63n(1e9))
}
