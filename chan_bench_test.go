package main

import "testing"
import "sync"

func BenchmarkChannels(b *testing.B) {
    ch := make(chan bool, 10000)

    waitGroup := sync.WaitGroup{}
    waitGroup.Add(2)

    go func() {
        for i := 0; i < b.N; i++ {
            ch <- true
        }
        waitGroup.Done()
    }()

    go func() {
        for i := 0; i < b.N; i++ {
            <- ch
        }
        waitGroup.Done()
    }()

    waitGroup.Wait()
}
