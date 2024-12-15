<?php

//pcntl_fork(), PHP'de bir işlem (process) oluşturmak için kullanılan bir fonksiyondur.
//Bu fonksiyon, mevcut çalışan işlemi (ana işlem veya parent process) kopyalar.
//ve yeni bir çocuk işlem (child process) oluşturur.

//Fork (çatallama), bir işlemden (process) bir kopya çıkarma işlemidir.
//Bu işlem, genellikle çoklu görev (multitasking) ve paralel işlem çalıştırmak için kullanılır.
//Çocuk işlem, ebeveyn işlemin (parent process) aynısıdır, ancak farklı bir işlem kimliğine (process ID - PID) sahiptir.
//Ebeveyn ve çocuk işlemler aynı kodu çalıştırır, ancak pcntl_fork()'un döndürdüğü değere göre farklı yollar izlenir:
//0 dönerse: Bu kod, çocuk işlemde (child process) çalıştırılmaktadır.
//Pozitif bir sayı dönerse: Bu kod, ebeveyn işlemde (parent process) çalıştırılmaktadır ve dönen sayı çocuk işlemin PID'sidir.
//-1 dönerse: Fork işlemi başarısız olmuştur.

//pcntl_fork() Avantajları
//Concurrency (Eşzamanlılık): İşlemleri paralel çalıştırarak zaman kazandırır.
//Verimlilik: CPU'nun tüm çekirdeklerini daha etkin kullanır.
//Esneklik: Ağ ve dosya işlemleri gibi yavaş işlemleri daha hızlı hale getirir.

//pcntl_fork() Dezavantajları
//Platform Kısıtlaması:
//Sadece Unix tabanlı sistemlerde çalışır (Linux, macOS). Windows'ta desteklenmez.
//Hafıza Kullanımı:
//Parent işlemin bir kopyası oluşturulduğu için daha fazla bellek kullanır.
//Zor Hata Yönetimi:
//Çoklu işlem çalıştırmak karmaşıklaşabilir ve deadlock veya yarış durumu (race condition) gibi sorunlar yaşanabilir.

$startTime = microtime(true);

$pid = pcntl_fork();

if ($pid == -1) {
    die("Fork failed!");
} elseif ($pid == 0) {
    $url = "https://example.com";
    $ch = curl_init($url);
    curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);
    curl_exec($ch);
    curl_close($ch);
    echo "Child process: Request sent.\n";
    exit(0);
} else {
    for ($i = 0; $i < 5; $i++) {
        echo "Parent process: $i\n";
        sleep(1); // Simulate workload
    }

    pcntl_wait($status); // Wait for the child process to finish
}

$endTime = microtime(true);

echo "Execution time: " . ($endTime - $startTime) . " seconds\n";

