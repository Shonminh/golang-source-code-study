package main

import (
	"fmt"
	"sync"
)

func main1() {
	a := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		for i := 0; i < 10; i++ {
			a <- i
		}
		close(a)
		wg.Done()
	}()

	go func() {
		for val := range a {
			fmt.Printf("num is %d\n", val)
		}
		fmt.Println("Channel closed")
		wg.Done()
	}()
	wg.Wait()
}

// 交替打印
func main() {
	a := []string{"1", "2", "3", "4", "5", "6"}
	b := []string{"A", "B", "C", "D", "E", "F"}
	wg := sync.WaitGroup{}
	wg.Add(2)
	c := make(chan int)
	d := make(chan int)
	go func() {
		for i := 0; i < 6; i += 2 {
			fmt.Printf("%s%s", a[i], a[i+1])
			c <- 1
			<-d
		}
		wg.Done()
	}()

	go func() {
		for i := 0; i < 6; i += 2 {
			<-c
			fmt.Printf("%s%s", b[i], b[i+1])
			d <- 1
		}
		wg.Done()
	}()

	wg.Wait()
}

// func main() {
// 	s := map[interface{}]bool{}
// 	s["sada"] = true
// 	s[1123] = true
// 	fmt.Println(s)
//
// 	var x interface{} = "sddas"
// 	var y interface{} = 213
// 	// z := x > y
// 	a := demo{}
// 	b := map[int]string{}
// 	c := a == demo{}
// 	fmt.Println(x == y)
// }

// type demo struct {
// 	A int
// }
//
// type demo1 struct {
// 	A int
// }

type Cat struct {
	a int
}

func (c *Cat) Quack() {
	fmt.Printf("a=%v", c.a)
	c.a++
}

type Duck interface {
	Quack()
}

func main3() {
	var d Duck = &Cat{}
	d.Quack()
	d.Quack()
}
