package funcdesc

import(
	"fmt"
)

func NoReturn()  {
	fmt.Println("no return !")
}

//两个参数一样类型，那就可以这样简约声明
func IntReturn(a ,b int) int  {
	return a+b
}
//多返回值，总额，平均数
func MulReturn(a,b,c int) (int, int)  {
	return a+b+c,(a+b+c)/3
}

//固定返回值的名称，返回的时候就不需要指定来。
func ReturnV(a,b,c int) (sum int, avg int)  {
	sum =a+b+c
	avg=(a+b+c)/3
	return  
}

//可变参数
func MulReturnB(c ...int) (sum int, avg int)  {
	sum =0
	avg =0
	for i := 0; i < len(c); i++ {
		sum+=c[i]
	}
	if len(c)>0{
		avg=sum/len(c)
	}
	return  
}


//可变参数
func MulReturnC(a int,c ...int) (sum int, avg int)  {
	sum =a
	avg =0
	for i := 0; i < len(c); i++ {
		sum+=c[i]
	}
	 
	avg=sum/(len(c)+1)
	 
	return  
}

//匿名函数使用
func Anonymous()  {
	fmt.Println("............... Anonymous start ...............")

	//afunc 是一个函数
	tmpfunc := IntReturn
	fmt.Printf("var func .tmpfunc type= %T \n",tmpfunc)
	fmt.Println("var func .tmpfunc(12,29)=",tmpfunc(12,29))

	afunc := func ( a,b int) int {
		return a-b
	}
	fmt.Printf("Anonymous.tmpfunc type= %T \n",afunc)
	fmt.Println("Anonymous.tmpfunc(12,29)=",afunc(12,29))


	var a int =199
	// defer 匿名函数使用调用方式
	defer func ()  {
		fmt.Printf("defer Anonymous.func is work a:%d!\n",a)
	}()

	a = 299
	fmt.Printf("last op is work a:%d!\n",a)

	
}


func add(a,b int) int  {
	return a+b
}
func sub(a,b int) int  {
	return a- b
}
func Varfunc(a,b int ,op func(int,int)int)  {
	fmt.Println("............... Varfunc start ...............")

	fmt.Printf("Varfunc.op(%d,%d)=%d \n",a,b,op(12,29))
}

//函数变量调用
func TestVarFunc()  {
	Varfunc(12,3,add)
	Varfunc(12,3,sub)
}


//闭包，一个函数与其相关的引用环境组合而成的实体
