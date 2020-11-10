package main

import (
	"fmt"
	"training003/adderex"
	"training003/deferex"
	"training003/example"
	"training003/funcdesc"
	"training003/varpart"
)

func main() {
	fmt.Println("............... training003 main start ...............")

	funcdesc.NoReturn()
	fmt.Println("funcdesc.IntReturn(12,29)=", funcdesc.IntReturn(12, 29))
	sum, avg := funcdesc.MulReturn(12, 34, 29)
	fmt.Printf("funcdesc.MulReturn(12,34,29),sum=%d ,avg=%d \n", sum, avg)

	sum1, avg1 := funcdesc.MulReturn(12, 34, 29)
	fmt.Printf("funcdesc.ReturnV(12,34,29),sum1=%d ,avg1=%d \n", sum1, avg1)
	//_ 下划线，忽略返回值。两个下划线就与没有使用变量接收返回值一样的。
	sum2, _ := funcdesc.MulReturn(12, 34, 29)
	fmt.Printf("funcdesc.ReturnV(12,34,29),sum3=%d   \n", sum2)

	sum3, avg3 := funcdesc.MulReturnB(12, 34, 29, 342, 342, 54, 345, 34)
	fmt.Printf("funcdesc.MulReturnB(12,34,29,342,342,54,345,34),sum3=%d ,avg3=%d   \n", sum3, avg3)

	sum4, avg4 := funcdesc.MulReturnC(12, 34, 29, 342, 342, 54, 345, 34)
	fmt.Printf("funcdesc.MulReturnC(12,34,29,342,342,54,345,34),sum4=%d ,avg4=%d   \n", sum4, avg4)

	deferex.Test_Defer()
	deferex.Test_Defer2()
	deferex.Test_Defer3()

	example.AllPrimeNumbers(100)
	fmt.Println("example.IsSXH(153)=", example.IsSXH(153))
	fmt.Println("example.IsSXH(152)=", example.IsSXH(152))

	str := "dlskj dwe得瑟23109 23  "
	fmt.Println(".CharCount：", str)

	num, en, space, other := example.CharCount(str)

	fmt.Printf(".CharCount：num:%d,en:%d,space:%d,other:%d \n", num, en, space, other)

	varpart.TestGlobalVar()
	varpart.TestLocalVar()

	funcdesc.Anonymous()
	funcdesc.TestVarFunc()

	adderex.TestAdder()

	fmt.Println("deferex.Defer1():", deferex.Defer1())
	fmt.Println("deferex.Defer2(11):", deferex.Defer2(11))

	deferex.TestReturnDeferMain()
	fmt.Println("............... training003 main end ...............")

}
