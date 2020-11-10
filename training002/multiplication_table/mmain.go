package multiplication_table

import (
	"fmt"
)


//乘法表
func Multiplication(num int)  {
	for i := 1; i <= num; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t",i,j,i*j)
		}
		fmt.Println()
	}
}