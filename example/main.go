package main

import (
	"context"
	"fmt"
	"time"
)

func main() {

	// ch := make(chan struct{}, 1)
	// ch <- struct{}{}
	// ch <- struct{}{}
	// // after := time.After(3 * time.Second)
	// // go func() {
	// select {
	// case <-ch:
	// 	fmt.Println("get")
	// 	// default:
	// }
	// }()
	// fmt.Println("asx")

	// background := context.Background()
	// go runWithTimeout(background)
	testAfterFunc()
	time.Sleep(10 * time.Second)
	// heap.Push()
}

func testAfterFunc() {
	time.AfterFunc(time.Second, func() {
		fmt.Println("schedule!")
	})
}

func runWithTimeout(background context.Context) {
	timeout, cancelFunc := context.WithTimeout(background, time.Second*2)
	defer cancelFunc()
	// time.Sleep(time.Second * 4)
	select {
	case <-timeout.Done():
		fmt.Println("sadfsdf")
	}
	// timeout.Done()
	// deadline, ok := timeout.Deadline()
	// fmt.Println(deadline, ok)
}
