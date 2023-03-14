package main

import (
	"fmt"
	"sync"
)

// 使用两个 goroutine 交替打印序列，
// 一个 goroutine 打印数字，
// 另外一个 goroutine 打印字母，
// 最终效果如下：
//
// 12AB34CD56EF

func main() {
	var A = []string{"1", "2", "3", "4", "5", "6"}
	var B = []string{"A", "B", "C", "D", "E", "F"}
	wg := sync.WaitGroup{}
	wg.Add(2)
	var c = make(chan int)
	var d = make(chan int)
	go func() {
		for i := 0; i < 6; i += 2 {
			fmt.Printf("%s%s", A[i], A[i+1])
			c <- 1
			<-d
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 6; i += 2 {
			<-c
			fmt.Printf("%s%s", B[i], B[i+1])
			d <- 1
		}
		wg.Done()
	}()
	wg.Wait()
}