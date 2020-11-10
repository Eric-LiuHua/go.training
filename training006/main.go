package main

import (
	"fmt"
	"strings"
	"training006/golable"
	"training006/structex"
	"training006/zzexample"

	"go.example/fjb"
)

func main() {

	fjb.Fjb(strings.Split("qroqeo,qi,alewp,era,qz,e,e,dxawhd", ","))
	zzexample.Testmain()

	golable.ForBreak()
	golable.ForSwithBreak()
	golable.ForSwithContinue()
	golable.GotoTag()

	var pp *structex.People = structex.NewPeople(1, "code1", "name1")
	fmt.Printf("base people:%v \n", pp.ToJson())

	structex.ChangeNameA(*pp)
	fmt.Printf("ChangeNameA outside people:%v \n", (*pp).ToJson())

	structex.ChangeNameB(pp)

	fmt.Printf("ChangeNameB outside people:%v \n", (*pp).ToJson())

	structex.TestAddrs()
}
