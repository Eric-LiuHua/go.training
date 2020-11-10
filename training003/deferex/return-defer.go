package deferex

import (
	"fmt"
)

//main
func TestReturnDeferMain() {

	fmt.Println("************* RunDefer **************")
	RunDefer()

	fmt.Println("************* RunReturnDefer **************")
	name1 := RunReturnDefer()
	fmt.Println("RunReturnDefer end Hello ", name1)

	fmt.Println("************* RunReturnDefer1 **************")

	name := RunReturnDefer1()
	fmt.Println("RunReturnDefer1 end Hello ", name)
}

//普通函数
func RunDefer() {
	name := "init name "
	defer CallRunDefer(name)
	name = "After RunDefer"

	fmt.Println("RunDefer Hello ", name)
}

//普通函数
func CallRunDefer(name string) {
	fmt.Println("CallRunDefer Hello ", name)
}

//defer，return和未命名的返回值
func RunReturnDefer() string {
	name := "init name"
	//未命名返回值，所以 defer 无法对返回的值进行操作的
	defer CallReturnRunDefer(&name)
	name = "After RunDefer"

	return name
}
func CallReturnRunDefer(name *string) {
	*name = "Change"
	fmt.Println("CallReturnRunDefer Hello ", *name)
}

//defer，return和未命名的返回值
func RunReturnDefer1() (x string) {
	name := "init name"
	x = name
	//未命名返回值，所以 defer 无法对返回的值进行操作的
	defer CallReturnRunDefer1(&x)
	defer CallReturnRunDefer2(x)
	x = "After RunDefer"

	return
}

func CallReturnRunDefer1(name *string) {
	fmt.Println("Befor CallReturnRunDefer1 Hello ", *name)
	*name = "Change 1"
	fmt.Println("After CallReturnRunDefer1 Hello ", *name)
}

//值传递，不影响返回值
func CallReturnRunDefer2(name string) {
	fmt.Println("Befor CallReturnRunDefer2 Hello ", name)
	name = "Change 2"
	fmt.Println("After CallReturnRunDefer2 Hello ", name)
}
