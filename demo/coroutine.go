// Package demo
package main

import (
	"fmt"
	"strconv"
	"time"
)

func myprint(num int) {
	for i := 0; i < 10; i++ {
		fmt.Println("i am coroutine " + strconv.Itoa(num) + ", echo " + strconv.Itoa(i))
	}
}

func main() {
	for j := 0; j < 10000; j++ {
		go myprint(j)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("i am	main!")
}
