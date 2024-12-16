package main

import (
	"fmt"
	"sync"
	"time"
)

func addition() int {
	// Simulating addition task
	time.Sleep(1 * time.Second)
	return 5 + 10
}

func multiplication() int {
	// Simulating multiplication task
	time.Sleep(1 * time.Second)
	return 5 * 10
}

func additionWg(wg *sync.WaitGroup) int {
	defer wg.Done()
	// Simulating addition task
	time.Sleep(1 * time.Second)
	return 5 + 10
}

func multiplicationWg(wg *sync.WaitGroup) int {
	defer wg.Done()
	// Simulating multiplication task
	time.Sleep(1 * time.Second)
	return 5 * 10
}

func mathConcurrently() {
	var wg sync.WaitGroup
	wg.Add(2) // We have two tasks to wait for

	// Concurrent execution: Both tasks start simultaneously
	go additionWg(&wg)
	go multiplicationWg(&wg)

	// Wait for both goroutines to finish
	wg.Wait()
}

func mathSequentially() {
	addition()
	multiplication()
}

func main() {
	start := time.Now()
	mathSequentially()
	fmt.Println("Execution Time:", time.Since(start))
}
