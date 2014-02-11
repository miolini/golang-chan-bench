package main

import "log"
import "time"

func main() {
    log.Printf("channel test")
    ch := make(chan bool, 10000)
    go func() {
        var count int
        var timer time.Time
        for {
            count = 0
            timer = time.Now()
            for {
                <- ch
                count++
                if count % 10000 == 0 && time.Since(timer) > time.Second * 5 {
                    break
                }
            }
            log.Printf("speed %dK msg/s", count/5/1000)
        }
    }()
    go func() {
        for {
            ch <- true
        }
    }()
    <- make(chan bool)
}
