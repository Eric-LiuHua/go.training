package forCycle

import (
	"fmt"
)


func For1(num int)  {
	fmt.Println("............... For1 start ...............")

	for i := 0; i < num; i++ {
		if i==9 {
			fmt.Printf("for i =%d,break \n",i)
			break
		}
		if i&1!=1{
			continue
		}

		fmt.Println("for i =",i)
	}
}

func For2(num int)  {
	fmt.Println("............... For2 start ...............")
	i:= 0
	for  ; i < num;   {
		fmt.Println("for i =",i)
		if i&1!=1{
			i+=2
		}else{
			i+=1
		}
	}
}

func For3(num int)  {
	fmt.Println("............... For3 start ...............")
	i:= 0
	for  i < num  {
		fmt.Println("for i =",i)
		if i&1!=1{
			i+=2
		}else{
			i+=1
		}
	}
}

 
//是否回文数
//多判断多for
func IsPalindrome(str string) bool{
	 
	var srune=[]rune(str)
	slen:=len(srune)

	if slen==0{
		return false
	}else if slen==1{
		return true
	}else{
		for i,j := 0,slen-1; i < j;i,j=i+1,j-1 {
			if srune[i]!=srune[j]{
				return false
			}
		}
		return true
	}
}