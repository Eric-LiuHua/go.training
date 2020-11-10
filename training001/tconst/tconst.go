package tconst

import (
	"fmt"
)


var(
	id int
	name string
	health bool
)
// 这种声明，都是默认跟上上一个
const (
	tmpid =100
	tmpid1 =200
	tmpid2 
)
 
// 默认从0开始递增
const (
	a = iota
	b  
	c 
)

//利用递增来进行二进制位移计算
const (
	a1 = 1<< iota
	b1 
	c1 
)
// 外部调用的函数名要求是首字母大写。
func Tmain()  {
	fmt.Printf("hello go,tmpid:%d,tmpid1:%d,tmpid2:%d \n",tmpid,tmpid1,tmpid2)
	fmt.Printf("hello const1,a:%d,b:%d,c:%d \n",a,b,c)
	fmt.Printf("hello const2,a1:%d,b1:%d,c1:%d \n",a1,b1,c1)
}