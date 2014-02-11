package main

import "testing"

func BenchmarkChannels(b *testing.B) {
    ch := make(chan bool, 10000)
    go func() {
        for _ = range ch {
        }
    }()
    for i := 0; i < b.N; i++ {
        ch <- true
    }
    close(ch)
}
