package main

import (
	"fmt"
	"training004/sorts"
)

func main() {
	fmt.Println("............... training004 main start ...............")
	b1 := []int{7, 89, 99, 55, 2235, 515, 88, 22, 2}
	fmt.Println("b1 org:", b1)
	b1 = sorts.Sort_Bubble(b1[:])
	fmt.Println("b1 org:", b1)

	b2 := []int{7, 2, 43, 35, 466, 23, 89, 99, 55, 2235, 515, 88, 22, 2}
	fmt.Println("b2 org:", b2)
	b2 = sorts.Sort_Counter(b2[:])
	fmt.Println("b2 org:", b2)

	b3 := []int{7, 2, 43, 35, 466, 23, 89, 99, 55, 2235, 515, 88, 22, 2}
	fmt.Println("b3 org:", b3)
	b3 = sorts.Sort_Merge(b3[:])
	fmt.Println("b3 org:", b3)

	fmt.Println("............... training004 main end ...............")
}
