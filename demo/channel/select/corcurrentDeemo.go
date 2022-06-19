package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

// go 并发编程
// 两个程序并发打印数据，一个打印小写字母，一个打印大写字母

// 用来生成消息
func genMsg(name string) <-chan string {
	c := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Duration(rand.Intn(2000)) * time.Millisecond)
			c <- fmt.Sprintf("service %s send data to channel %s", name, strconv.Itoa(i))
		}
	}()
	return c
}

func printMsgCorcurrent(m1 <-chan string, m2 <-chan string, wg *sync.WaitGroup) {
	go func() {
		for {
			fmt.Println(<-m1)
			wg.Done()
		}
	}()

	go func() {
		for {
			fmt.Println(<-m2)
			wg.Done()
		}
	}()
}

func fanIn(m1, m2 <-chan string) chan string {
	m := make(chan string)
	go func() {
		for {
			m <- <-m1
		}
	}()

	go func() {
		for {
			m <- <-m2
		}
	}()

	return m
}
func printMsgBySelect(m1, m2 <-chan string) {
}

func main() {
	m1 := genMsg("service1")
	m2 := genMsg("service2")

	//printMsgCorcurrent(m1, m2)

	//// 下面两个for循环，虽然可以正常接收数据，但不是并发的，需要一个channel
	//// 读完才去读另外一个channel, 可以进一步改造成并发的，使用select
	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-m1)
	//}

	//for i := 0; i < 10; i++ {
	//	fmt.Println(<-m2)
	//}

	m := fanIn(m1, m2)
	for i := 0; i < 20; i++ {
		fmt.Println(<-m)
	}
	//time.Sleep(2 * time.Second)
	fmt.Println("select")
}
