<?php

$startTime = microtime(true);

$url = "https://example.com";
$ch = curl_init($url);
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
curl_exec($ch);
curl_close($ch);

echo "Synchronous: Request sent.\n";

for ($i = 0; $i < 5; $i++) {
    echo "Synchronous: $i\n";
    sleep(1); // Simulate workload
}

$endTime = microtime(true);

echo "Execution time: " . ($endTime - $startTime) . " seconds\n";



