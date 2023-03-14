package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main1() {

	var total int64
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		go func() {
			wg.Add(1)
			for true {
				var old = total
				addInt64 := atomic.AddInt64(&total, 1)
				if old == addInt64 {
					fmt.Println("error")
				} else {
					fmt.Printf("old=%v, new=%v\n", old, addInt64)
				}
			}
		}()
	}
	time.Sleep(time.Second)
	wg.Wait()
}
