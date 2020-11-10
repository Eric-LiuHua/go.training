package process_control


import (
	"fmt"
)

// is 条件分支
func Testcondition(v int)  {
	if v<=0 {
		fmt.Println("entry if v<=0 ,v=",v)
	}else if v <=10{
		fmt.Println("entry else if v <=10 ,v=",v)
	}else{
		fmt.Println("entry else ,v=",v)
	}
}

//是否偶数
func IsEvenNumber(num int) {
	if num&1 !=1 {
		fmt.Println(num," is even number.")
	}else{
		fmt.Println(num," is not even number.")
	}
}

func IsEvenNumber1() {
	if num:=GetNum(); num&1 !=1 {
		fmt.Printf("%d is even number.\n",num)
	}else{
		fmt.Printf("%d is not even number. \n",GetNum())
	}
}

func GetNum() int {
	return 100
}