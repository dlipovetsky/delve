package main

import (
	"fmt"
	"time"
)

func main() {
	// Number of goroutines to run.
	n := 3

	done := make(chan interface{})

	for i := 0; i < n; i++ {
		go func(j int, done <-chan interface{}) {
			for {
				select {
				case <-done:
					return
				case <-time.After(1 * time.Second):
					fmt.Printf("goroutine %d\n", j)
				}
			}
		}(i, done)
	}

	time.Sleep(1 * time.Hour)
	close(done)
}
