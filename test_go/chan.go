package main

import (
    "fmt"
    "sync"
    "time"
)

func main() {
    testNoBufferChan()
    testBufferChan()
}

func testNoBufferChan() {
    wg := new(sync.WaitGroup)
    ch := make(chan bool)

    wg.Add(1)
    go func() {
        defer wg.Done()

        time.Sleep(10 * time.Second)

        fmt.Println(time.Now(), "NoBufferChan recv begin")
        <-ch
        fmt.Println(time.Now(), "NoBufferChan recv end")
    }()

    fmt.Println(time.Now(), "NoBufferChan send begin")
    ch <- true
    fmt.Println(time.Now(), "NoBufferChan send end")

    wg.Wait()
}

func testBufferChan() {
    wg := new(sync.WaitGroup)
    ch := make(chan bool, 1)

    wg.Add(1)
    go func() {
        defer wg.Done()

        time.Sleep(10 * time.Second)

        fmt.Println(time.Now(), "BufferChan recv begin")
        <-ch
        fmt.Println(time.Now(), "BufferChan recv end")
    }()

    fmt.Println(time.Now(), "BufferChan send begin")
    ch <- true
    fmt.Println(time.Now(), "BufferChan send end")

    wg.Wait()
}
