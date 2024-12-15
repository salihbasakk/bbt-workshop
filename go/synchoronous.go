package main

import (
	"fmt"
	"time"
	"runtime"
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

func main() {
	// Track execution time
	start := time.Now()

	// Memory usage at the start
	var mStart runtime.MemStats
	runtime.ReadMemStats(&mStart)

	// Synchronous execution: One after the other
	sum := addition()
	multiply := multiplication()

	// Print results
	fmt.Printf("Addition result: %d\n", sum)
	fmt.Printf("Multiplication result: %d\n", multiply)

	// Memory usage at the end
	var mEnd runtime.MemStats
	runtime.ReadMemStats(&mEnd)

	// Execution time
	fmt.Printf("Execution time: %.6fs\n", time.Since(start).Seconds())

	// Memory usage difference
	fmt.Printf("Memory used: %d bytes\n", mEnd.Alloc-mStart.Alloc)
}