package varpart

import(
	"fmt"
)

//全局变量
var globalv int  =9999

func TestGlobalVar()  {
	fmt.Println("全局作用域 TestGlobalVar:",globalv)
}

//全局
func TestLocalVar()  {
	var localvar int =91 
	fmt.Println("变量的优先级，如果重名，优先使用最小作用域的函数作用域 globalv:",globalv)

	var globalv int =91 
	fmt.Println("函数作用域 localvar:",localvar)
	fmt.Println("变量的优先级，如果重名，优先使用最小作用域的 globalv:",globalv)

	for i := 0; i < 5; i++ {
		tmp:= i* localvar
		fmt.Println("方法块作用域 tmp:",tmp)

		if globalv:=i ;globalv>3 {
			fmt.Println("方法块作用域globalv>3 globalv:",globalv)
		}else
		{
			fmt.Println("方法块作用域globalv<=3 globalv:",globalv)
		}
	}
	

}