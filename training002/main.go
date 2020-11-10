package main

import(
	"fmt"
	 "training002/process_control"
	 "training002/forCycle"
	 "training002/switch_control"
	 "training002/multiplication_table"
	 
)
//多变量赋值
func margs()  {
	var a int
	var b int
	a,b =32,33

	fmt.Printf("margs -> a:%d,b:%d \n",a,b)
}

func main()  {
	fmt.Println("............... training main start ...............")
	process_control.IsEvenNumber(-11)
	process_control.IsEvenNumber(2)
	process_control.IsEvenNumber(231)
	process_control.IsEvenNumber1()

	forCycle.For1(11)
	forCycle.For2(13)
	forCycle.For3(13)

	margs()

	switch_control.TestSwitch(3244)
	switch_control.TestSwitch(2223)
	switch_control.TestSwitch2(424)

	multiplication_table.Multiplication(9)

	fmt.Println("............... training main end ...............")
}