package main

import (
	"fmt"
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

func main() {
	// Track execution time
	start := time.Now()

	// Synchronous execution: One after the other
	addition()
	multiplication()

	// Execution time
	fmt.Printf("Execution time: %.6fs\n", time.Since(start).Seconds())
}