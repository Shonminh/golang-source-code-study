package main

import (
	"fmt"
)

func demo1() {
	var a, b int

}

func printNow(n int) {
	c := n / 2
	for i := 0; i <= c; i++ {
		for j := i; j < c; j++ {
			fmt.Printf(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
	for i := c - 1; i >= 0; i-- {
		for j := c; j > i; j-- {
			fmt.Printf(" ")
		}
		for j := 0; j < 2*i+1; j++ {
			fmt.Printf("*")
		}
		fmt.Println()
	}
}

func main() {
	printNow(5)
	printNow(3)
	printNow(10)
	// a := "sadfff:"
	// split := strings.Split(a, ":")
	var a []int
	a = append(a, 123)
}

// netstat -an | grep TIME_WAIT | awk '{print $5}' | awk -F ':' '{print $1}' | sort | uniq -c | sort -nr
