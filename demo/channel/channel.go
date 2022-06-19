package main

import (
	"fmt"
	"time"
)

func worker(num int, c chan int) {
	for {
		fmt.Printf("worker	%d reciew %c\n", num, <-c)
	}
}

func chanDemo() {
	var channels [10]chan int
	for i := 0; i < 10; i++ {
		channels[i] = make(chan int)
		go worker(i, channels[i])
	}

	// 向channel中发数据
	for j := 0; j < 10; j++ {
		channels[j] <- 'a' + j
	}

	for j := 0; j < 10; j++ {
		channels[j] <- 'A' + j
	}
	time.Sleep(time.Millisecond)
}

func main() {

	chanDemo()
	fmt.Println("main exit")
}
