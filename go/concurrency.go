package main

import (
	"fmt"
	"time"
	"runtime"
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

	// Memory usage at the start
	var mStart runtime.MemStats
	runtime.ReadMemStats(&mStart)

	// Use WaitGroup for synchronization of goroutines
	var wg sync.WaitGroup
	wg.Add(2) // We have two tasks to wait for

	// Concurrent execution: Both tasks start simultaneously
	go addition(&wg)
	go multiplication(&wg)

	// Wait for both goroutines to finish
	wg.Wait()

	// Memory usage at the end
	var mEnd runtime.MemStats
	runtime.ReadMemStats(&mEnd)

	// Execution time
	fmt.Printf("Execution time: %.6fs\n", time.Since(start).Seconds())

	// Memory usage difference
	fmt.Printf("Memory used: %d bytes\n", mEnd.Alloc-mStart.Alloc)
}