package deferex

import "fmt"

//Defer1 2 x:1  ,看看出，在return之前先倒序调用defer ，
//Defer1 1 x:2 看到倒序第二个是，x++，已经对变量起了作用，
//deferex.Defer1(): 0 ，虽然先执行了defer ，但是不会影响return对结果
func Defer1() int {
	var x int
	defer func() {
		x++
		fmt.Printf("Defer1 1 x:%d \n", x)
	}()

	defer func() {
		x++
		fmt.Printf("Defer1 2 x:%d \n", x)
	}()

	return x
}

//Defer2 2 x:12
//Defer2 1 x:13
//deferex.Defer2(11): 11
func Defer2(x int) int {

	defer func() {
		x++
		fmt.Printf("defer 1 x:%d \n", x)
	}()

	defer func() {
		x++
		fmt.Printf("defer 2 x:%d \n", x)
	}()

	return x
}
