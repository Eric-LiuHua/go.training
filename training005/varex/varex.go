package varex

import "fmt"

//每个变量都有内存地址。变量的赋值和修改，都是通过修改对应内存地址
func VarTest1() {
	var a int
	a = 100
	fmt.Printf("a=%d ,the addr of a :%p \n", a, &a)
}

//*int 是，int 类型的指针。
func change(v *int, tmp int) {
	//对应内存的值改为tmp
	*v = tmp
}

//值拷贝
func changev(v int) {
	v = 999
}

//修改数组指针的值,需要强转类型，才可以赋值
func ModifyArr(arr *[3]int) {
	if len(*arr) > 0 {
		(*arr)[0] = 111
	}
}

//todo & 取变量的地址，*取地址的值
func Pointer() {
	a := 58
	fmt.Println("1.a=", a)
	fmt.Println("1.a=", &a)
	b := &a
	c := a
	changev(c)
	fmt.Println("changev(c) 值拷贝 ")
	fmt.Println("2.a=", a)
	fmt.Println("2.指针传参.b=", *b)
	fmt.Println("2.值传递.c=", c)

	*b = 1233
	fmt.Println("修改*b= 1233")
	fmt.Println("3.b=", *b)
	fmt.Println("3.a=", a)
	fmt.Println("3.c=", c)

	fmt.Println("4.a=", &a)
	fmt.Println("4.b=", b)
	fmt.Println("4.c=", &c)

}

func ArrPointer() {
	arr := [3]int{32, 43, 55}
	fmt.Println("ArrPointer befor arr=", arr)
	//传递arr 的内存地址
	ModifyArr(&arr)

	fmt.Println("ArrPointer after arr=", arr)
}

//切片是引用类型，切片存的是地址
func ModifySlice(slice []int) {
	slice[0] = 9239
}

func SlicePointer() {
	arr := [3]int{22, 32, 434}
	fmt.Println("SlicePointer befor arr=", arr)
	sliceArr := arr[:]
	ModifySlice(sliceArr)
	fmt.Println("SlicePointer after arr=", arr)

}

func Pointer1() {
	var a int = 99
	var ip *int

	fmt.Printf("空指针 ip 的值为 : %x ,%v \n", ip, (ip == nil))
	//把a的地址给指针
	ip = &a
	fmt.Printf("a 变量的地址是: %x \n", &a)

	/* 指针变量的存储地址 */
	fmt.Printf("ip 变量储存的指针地址: %x\n", ip)

	/* 使用指针访问值 */
	fmt.Printf("*ip 变量的值: %d\n", *ip)
}
