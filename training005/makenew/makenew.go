package makenew

import "fmt"

//make 用来分配引用类型的内存，比如map，slice，channel
//new 用来分配除引用类型之外所以其他类型的的内存，比如int，数组等

func TestPointer1() {
	//a 指向一个int 大小的内存地址。在给地址赋值
	var a *int = new(int)
	//操作a指向那块内存
	*a = 99
	fmt.Println("*a= ", *a)

	//分配一个指针指向一个空切片的地址。
	var b *[]int = new([]int)
	fmt.Printf("after new *b= %v \n", *b)

	//(*b)[0]=99  //不能直接操作这个空切片。index out of range

	(*b) = make([]int, 10, 10)
	(*b)[0]=99
	(*b)[2]=99
	(*b)[3]=99
	fmt.Printf("after make  *b= %v \n", *b)

}
