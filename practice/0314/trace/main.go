package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
	"time"
)

func main() {
	file, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	err = trace.Start(file)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	wg := sync.WaitGroup{}
	wg.Add(1)
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			<-ch
			fmt.Printf("i=%v\n", i)
		}
		wg.Done()
	}()

	for i := 0; i < 3; i++ {
		ch <- 1
		time.Sleep(time.Second)
	}
	wg.Wait()
}
