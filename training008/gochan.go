package main

import (
	"fmt"
)

func TestSimpleChan() {
	//0 缓冲通道，会堵塞协程。那协程对a对改变之后，
	//>0 有缓冲对通道。那就是异步来。除非缓冲满来。否则是不堵塞的
	var c = make(chan int, 0)
	var a string
	go func() {
		a = "dxxdf"
		<-c
	}()
	c <- 1
	fmt.Println("SimpleChan:", a)
}

//关闭通道
func TestCloseChan() {
	var ch = make(chan int, 3)
	go func() {
		ch <- 1
		close(ch)
	}()

	for cm := range ch {
		fmt.Println("TestCloseChan,cm:", cm)
	}
	//因为已经关闭来。这里取到默认值，0，表明通道已经没有数据来
	fmt.Println(<-ch)
}
