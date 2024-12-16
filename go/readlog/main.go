package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
	"time"
)

const logFilePath = "go/readlog/app.log"
const logPrefix = "LOG:"

func processLogLine(line string) {
	time.Sleep(500 * time.Millisecond)
	fmt.Println(logPrefix, line)
}

func processLogLineWg(line string, wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println(logPrefix, line)
}

func readLogSequentially(filePath string) {
	// File Open: Read Only
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	defer file.Close()

	// Read Line: Sequentially (One after the other)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		processLogLine(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Err:", err)
	}
}

func readLogConcurrently(filePath string) {
	// File Open: Read Only
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Err:", err)
		return
	}
	defer file.Close()

	var wg sync.WaitGroup

	// Read Line: Concurrently (Multiple at the same time)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		wg.Add(1) // Yeni bir iş ekle
		go processLogLineWg(line, &wg)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Err:", err)
	}

	wg.Wait() // Wait for all goroutines to finish
}

func main() {
	start := time.Now()
	readLogConcurrently(logFilePath)
	fmt.Println("Execution Time:", time.Since(start))
}
