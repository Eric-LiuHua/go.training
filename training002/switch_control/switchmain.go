package switch_control

import (
	"fmt"
	"time"
	"math/rand"
)


func TestSwitch(rnum int)  {
	fmt.Println("............... TestSwitch start ...............")
	num:=GetNum(rnum)
	switch num {
	case 1:
		fmt.Printf("case 1: switch i= %d,break \n",num)
	case 2:
		fmt.Printf("case 2: switch i= %d,break \n",num)
	case 3,8,4,9,5,6,7,11:
		fmt.Printf("case 3,8,4,9,5,6,7,11: switch i= %d,break \n",num)
	default:
		fmt.Printf("switch i= %d  not in(1,2,3),break \n",num)
	}
}

func TestSwitch2(rnum int)  {
	fmt.Println("............... TestSwitch2 start ...............")

	num:=GetNum(rnum)
	switch {
	case num >1 && num <=10:
		fmt.Printf("switch i =%d num <=0,break \n",num)
	case num>10 && num <222:
		fmt.Printf("switch i =%d num>0 && num <12,break \n",num)
	case num>=1223:
		fmt.Printf("switch i =%d num>=12 ,break \n",num)
	default:
		fmt.Printf("switch i= %d not in(1,2,3),break \n",num)
	}
}

func GetNum(rnum int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(rnum)
}