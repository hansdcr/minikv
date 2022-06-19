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

func fanInSelect(c1, c2 <-chan string) chan string {
	c := make(chan string)
	go func() {
		for {
			select {
			case m := <-c1:
				c <- m
			case m := <-c2:
				c <- m
			}
		}
	}()
	return c
}

func main() {
	c1 := genMsg("service1")
	c2 := genMsg("service2")

	//m := fanIn(c1, c2)
	c := fanInSelect(c1, c2)
	for i := 0; i < 20; i++ {
		fmt.Println(<-c)
	}
	fmt.Println("select")
}
