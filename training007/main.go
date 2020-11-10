package main

import (
	"fmt"
	"training007/osx"
	"training007/reflects"
)

func main() {
	fmt.Println("............... training007 main start ...............")

	reflects.TestReflect()

	osx.TestOsx()
	fmt.Println("............... training007 main end ...............")

}
