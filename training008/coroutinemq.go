package main

import (
	"fmt"
	"time"
)

//只写通道
func Producer(ch chan<- int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
}

//只读通道
func Consumer1(ch <-chan int) {
	for cm := range ch {
		fmt.Printf("Consumer1 hit num:%d \n", cm)
	}
}

//只读通道
func Consumer2(ch <-chan int) {
	for cm := range ch {
		fmt.Printf("Consumer2 hit num:%d \n", cm)
	}
}

//测试入口
func TestCoroutine() {
	var ch = make(chan int, 2)
	go Consumer1(ch)
	go Consumer2(ch)
	Producer(ch)
	time.Sleep(time.Second)
}
