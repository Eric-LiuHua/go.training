package adderex

import (
	"fmt"
	"strings"
	"time"
)

//闭包
func Adder(init int) func(a, b int) int {
	var x int = init
	return func(a, b int) int {
		x = a + b + x
		return x
	}
}

//处理后缀的闭包
func MakeSuffix(suffix string) func(string) string {
	return func(s string) string {
		if !strings.HasSuffix(s, suffix) {
			return s + suffix
		} else {
			return s
		}
	}
}

//闭包3 多返回函数形式
func Adder2(init int) (func(a, b int) int, func(a, b int) int) {
	var x int = init
	c1 := func(a, b int) int {
		x = a + b + x
		return x
	}
	c2 := func(a, b int) int {
		x = a + b - x
		return x
	}

	return c1, c2
}

//闭包的弊端，循环中调用外部变量一直在变,虽然传进去，但是顺序不一定
func Adder3(init int) {
	for i := 0; i < init; i++ {
		fmt.Printf("xx  %d \n", i)
		go func() { fmt.Printf("go func %d \n", i) }()
		go func(int2 int) { fmt.Printf("int2 go func %d \n", int2) }(i)
	}
	//没有睡眠的话，闭包不能正常执行的。可能不会被执行
	//time.Sleep(time.Second*2)
}

//TestAdder 测试闭包入口
func TestAdder() {
	time.Sleep(time.Second * 2)
	adder := Adder(90)

	fmt.Printf("adder(2,3):%d \n", adder(2, 3))
	fmt.Printf("adder(3,33):%d \n", adder(3, 33))

	suffix := MakeSuffix(".jpg")
	fmt.Printf("suffix(\"caonima\") = %s \n", suffix("caonima"))
	fmt.Printf("suffix(\"cccc.jps\") = %s \n", suffix("cccc.jps"))
	fmt.Printf("suffix(\"cccc.jpg\") = %s \n", suffix("cccc.jpg"))

	c1, c2 := Adder2(100)
	fmt.Printf("c1(2,3):%d,c1(2,2):%d \n", c1(2, 3), c2(2, 2))
	fmt.Printf("c1(2,3):%d,c1(2,2):%d \n", c1(33, 233), c2(872, 72))

	Adder3(3)
}
