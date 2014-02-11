package main

import "testing"
import "sync"

func BenchmarkChannels(b *testing.B) {
    ch := make(chan bool, 10000)

    var done sync.WaitGroup
    done.Add(1)

    go func() {
        for _ = range ch {
        }
        done.Done()
    }()

    for i := 0; i < b.N; i++ {
        ch <- true
    }
    close(ch)
    done.Wait()
}