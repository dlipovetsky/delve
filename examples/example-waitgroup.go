package main

import (
	"fmt"
	"sync"
)

func main() {
	n := 3 
	wg := sync.WaitGroup{}

	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(j int) {
			fmt.Printf("goroutine %d\n", j)
			defer wg.Done()
		}(i)
	}

	wg.Wait()
}
