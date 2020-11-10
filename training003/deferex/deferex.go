package deferex

import(
	"fmt"
)

//defer 测试
//常用释放资源，在函数结束之前执行。
//多个defer 方法栈的执行顺序。
func Test_Defer()  {
	defer fmt.Println(" defer Test defer work aa1!")
	defer fmt.Println(" defer Test defer work aa2!")
	defer fmt.Println(" defer Test defer work aa3!")
	fmt.Println(" Test defer work bb!")
	fmt.Println(" Test defer work cc!")
}

//for defer 测试
func Test_Defer2()  {
	for i := 0; i < 10; i++ {
		defer fmt.Printf(" defer Test_Defer2 work %d! \n",i)
	}

	defer fmt.Println(" defer Test_Defer2  work aa2!")
	 
 
	fmt.Println(" Test_Defer2  work cc!")
}

func Test_Defer3()  {
	i :=-1
	//defer 定义的时候，i 的值已经传入来。后续修改不会影响。
	defer fmt.Printf("defer Test_Defer3 work %d! \n",i)
 
 	i=1999
	fmt.Printf(" Test_Defer3  work aa2 ,%d! \n",i)

}