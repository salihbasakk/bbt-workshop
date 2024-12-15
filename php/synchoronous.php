<?php

$startTime = microtime(true);
$startMemory = memory_get_usage();  // Initial memory usage

// First process (Addition)
$sum = 5 + 10;
echo "Addition result: $sum\n";

// Second process (Multiplication)
$multiply = 5 * 10;
echo "Multiplication result: $multiply\n";

$endTime = microtime(true);
$endMemory = memory_get_usage();  // Final memory usage

echo "Execution time: " . ($endTime - $startTime) . " second\n";
echo "Memory used: " . ($endMemory - $startMemory) . " bytes\n";

//Concurrency kullandığınızda, her işlem (örneğin, toplama ve çarpma) ayrı bir işlem olarak çalıştırılır.
//Bu da her işlemin kendi belleği ve kaynaklarını talep etmesine neden olur.
//Senkron bir işlemde ise, işlemler sırasıyla tek bir iş parçacığında yapılır,
//dolayısıyla aynı anda birden fazla işlem yapılmaz ve bellek ve işlem kaynakları daha verimli kullanılır.

//Concurrency kullanmak, her işlem için yeni bir işlem başlatılması ve işlem yönetimi gerektirir.
//Bu da daha fazla bellek tüketimi ve daha fazla işlem süresi ile sonuçlanabilir.
//Senkron bir yaklaşımda ise, her işlem ardışık olarak çalıştırıldığı için daha düşük bellek tüketimi
//ve daha hızlı çalışma süresi elde edilir. Ancak, bu daha az ölçeklenebilir ve paralel işlem yapma yeteneğini kısıtlar.