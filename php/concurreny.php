<?php

$startTime = microtime(true);
$startMemory = memory_get_usage();  // Initial memory usage

// First process (Addition)
$pid1 = pcntl_fork();
if ($pid1 == -1) {
    die("Fork failed for process 1!");
} elseif ($pid1 == 0) {
    // Child process 1: Addition
    $sum = 5 + 10;
    echo "Addition result: $sum\n";
    exit(0);  // Exit child process
}

// Second process (Multiplication)
$pid2 = pcntl_fork();
if ($pid2 == -1) {
    die("Fork failed for process 2!");
} elseif ($pid2 == 0) {
    // Child process 2: Multiplication
    $multiply = 5 * 10;
    echo "Multiplication result: $multiply\n";
    exit(0);  // Exit child process
}

// Parent process waits for children to finish
pcntl_wait($status); // Wait for child 1 to finish
pcntl_wait($status); // Wait for child 2 to finish

$endTime = microtime(true);
$endMemory = memory_get_usage();  // Final memory usage

echo "Execution time: " . ($endTime - $startTime) . " second\n";
echo "Memory used: " . ($endMemory - $startMemory) . " bytes\n";  // Memory usage difference
