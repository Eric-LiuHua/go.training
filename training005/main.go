package main

import (
	"ex001"
	"fmt"
	"training005/arrayex"
	"training005/initex"
	"training005/mapex"
	"training005/varex"
)

func main() {
	arrayex.ArrayMain()
	mapex.MapTest4()
	varex.Pointer1()
	//fmt.Println("*************** var test ****************")
	//varex.VarTest1()

	//varex.Pointer()
	//varex.ArrPointer()

	//varex.SlicePointer()
	//makenew.TestPointer1()

	//调用静态台
	fmt.Println("dsdfsf:", ex001.Adds(12, 3, 234, 24, 3))

	initex.Tmp()
}
