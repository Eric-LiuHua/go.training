package main

import (
	"fmt"
	"time"
)

//简单的协程
func TestSimpleCoroutine() {
	fmt.Println("Simple coroutine is run .")
	for i := 0; i < 10; i++ {
		go func(i int) {
			fmt.Println("coroutine:", i)
		}(i)
	}
	//防止子协程还没有结束主协程就退出了
	time.Sleep(time.Second)
}
