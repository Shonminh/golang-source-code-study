package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"sync/atomic"
	"time"
)

func main1() {
	ctx := context.Background()
	ctx, cancelFunc := context.WithTimeout(ctx, time.Second)
	cancelFunc()
}
func main2() {
	// mu := sync.Mutex{}
	// mu.Lock()
	// mu.Unlock()
	// a := Demo{
	// 	struct{ A int }{A: 9},
	// 	struct{ A int }{A: 1},
	// }
	// fmt.Println(a)
	mu := sync.RWMutex{}
	mu.Lock()
	const rwmutexMaxReaders = 3
	for i := 0; i < rwmutexMaxReaders; i++ {
		mu.RLock()
	}
	mu.RUnlock()
	fmt.Println(mu)
}

type Demo [8]struct {
	A int
}

func main() {
	var status int64
	cond := sync.NewCond(&sync.Mutex{})
	for i := 0; i < 10; i++ {
		go func(cond *sync.Cond, i int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			for atomic.LoadInt64(&status) != 1 {
				cond.Wait()
			}
			fmt.Println("haha, status=", status)
		}(cond, i)
	}

	time.Sleep(time.Second)
	go func(cond *sync.Cond) {
		cond.L.Lock()
		defer cond.L.Unlock()
		atomic.AddInt64(&status, 1)
		for true {
			fmt.Println("asdfsadf")
			time.Sleep(time.Second)
		}
		cond.Broadcast()
	}(cond)
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
