package main

import (
	"fmt"
	"time"
	"sync"
)

func addition(wg *sync.WaitGroup) int {
	defer wg.Done()
	// Simulating addition task
	time.Sleep(1 * time.Second)
	return 5 + 10
}

func multiplication(wg *sync.WaitGroup) int {
	defer wg.Done()
	// Simulating multiplication task
	time.Sleep(1 * time.Second)
	return 5 * 10
}

func main() {
	// Track execution time
	start := time.Now()

	// Use WaitGroup for synchronization of goroutines
	var wg sync.WaitGroup
	wg.Add(2) // We have two tasks to wait for

	// Concurrent execution: Both tasks start simultaneously
	go addition(&wg)
	go multiplication(&wg)

	// Wait for both goroutines to finish
	wg.Wait()

	// Execution time
	fmt.Printf("Execution time: %.6fs\n", time.Since(start).Seconds())
}