package main

import (
	"fmt"
	"time"
	"training001/tconst"
	"training001/tstring"
	"training001/ttime"
	"training001/ttype"
)

// printf("%c",a)；输出单个字符。
// printf("%d",a)；输出十进制整数。
// printf("%f",a)；输出十进制浮点数.
// printf("%o",a)；输出八进制数。
// printf("%s",a)；输出字符串。
// printf("%u",a)；输出无符号十进制数。
// printf("%x",a)；输出十六进制数。
func main() {
	fmt.Println("............... training main start ...............")
	// 外部调用的函数名要求是首字母大写。
	tconst.Tmain()

	ttype.TestInt8()
	ttype.TestString()
	ttype.TestSlice()
	ttype.TestStrings()
	ttype.TestFloat64()

	tstring.TestString()
	tstring.TestRune()
	tstring.TestEqual()

	var str string = "测试字符串逆序，hello goo"
	fmt.Println("(str):", str)
	fmt.Println("tstring.RreverseString(str):", tstring.RreverseString(str))
	fmt.Println("tstring.IsPalindrome(str):", tstring.IsPalindrome(str))
	fmt.Println("tstring.IsPalindrome(\"nin\"):", tstring.IsPalindrome("nin"))

	ttime.TestTime()

	//定时器，
	ch := ttime.TestTicker()
	time.Sleep(5 * time.Second)
	ch <- true

	ttime.TestConst()
	ttime.TestFormat()

	ttime.TestCost()
	fmt.Println("............... training main end ...............")
}
