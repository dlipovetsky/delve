package main

import (
	"fmt"
	"sync"
)

func main() {
	// Number of goroutines to run.
	n := 3
	// Used to wait start goroutines in sequence, so that all are guaranteed to run.
	wg := sync.WaitGroup{}
	// Used to keep goroutines running until a signal is broadcast.
	c := sync.NewCond(&sync.Mutex{})

	c.L.Lock()
	for i := 0; i < n; i++ {
		wg.Add(1)
		go func(j int) {
			wg.Done()

			fmt.Printf("goroutine %d\n", j)

			c.L.Lock()
			c.Wait()
			c.L.Unlock()
		}(i)
		wg.Wait()
	}

	c.Broadcast()
}
