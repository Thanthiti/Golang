package main

import (
    "fmt"
    "sync"
    "time"
)

func downloadFile(name string, wg *sync.WaitGroup) {
    defer wg.Done()
    fmt.Printf("Downloading %s...\n", name)
    time.Sleep(2 * time.Second)
    fmt.Printf("Downloaded %s ✅\n", name)
}

func main() {
    files := []string{"fileA", "fileB", "fileC"}
    var wg sync.WaitGroup

    start := time.Now()

    for _, file := range files {
        wg.Add(1)
        go downloadFile(file, &wg)
    }

    wg.Wait() // รอให้ทุก goroutine เสร็จ

    elapsed := time.Since(start)
    fmt.Printf("Total time: %s\n", elapsed)
}
