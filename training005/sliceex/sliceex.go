package sliceex

import (
	"fmt"
)

//todo 切片的定义
//切片是基于数组类型做了一层封装，非常灵活，可以自动扩容
//切片基本操作：
//a.arr[start:end]包含start到end-1的所有元素切片
//b.arr[start:] 包含start 到最后一个元素的切片
//c.arr[:end] 包含0 到end-1 的所有元素切片
//d.arr[:]整个数组的切片
func slice() {
	var a []int = []int{212, 231}
	var c []int = []int{2122, 2321}
	var b [5]int = [5]int{1: 32, 2: 323}

	fmt.Println(a)
	fmt.Println(append(a, c...))
	fmt.Println(b)

}

func SliceInit() {
	fmt.Println("*************** SliceInit *********************")

	//定义int类型空切片
	var ea []int
	var b [5]int = [5]int{1: 32, 2: 323}
	if ea == nil {
		fmt.Println("ea is nil !")
	} else {
		fmt.Println("ea = ", ea)
	}
	//基于数组创建一个切片,b[start:end],得到start 到 end -1 的数据的切片
	ea = b[1:3]
	fmt.Println("b[1:3] = ", b[1:3])
	fmt.Println("b[2:] 0= ", b[2:])
	fmt.Println("b[:2] 1 = ", b[:2])
	fmt.Println("b[:] 2 = ", b[:])

	//通过数组创建切片
	eb := []int{12, 2, 3, 15, 89, 16, 13, 52}
	fmt.Println("eb = ", eb)

	//创建数组，其中[...] 是编译器的确定数组长度，有长度的是数组，没有长度的就是切片
	arr := [...]int{12, 14, 25, 3, 6, 7, 5, 85, 6, 8}

	ec := arr[2:]

	fmt.Println("befor ec:", ec)
	for i := range ec {
		ec[i]++
	}
	fmt.Println("after ec:", ec)

}

//切片就是数组的引用。修改切片的值可以改变数组的值
//切片和数组对值修改，都会引起其他的引用的一起变化
func SliceExample() {
	fmt.Println("*************** SliceExample *********************")
	arr := [5]int{2, 3, 34, 231, 23}
	s1 := arr[:]
	s2 := arr[:]
	fmt.Println("befor arr:", arr)
	fmt.Println("befor s1:", s1)
	fmt.Println("befor s2:", s2)
	s1[0] += 32
	fmt.Println("after s1[0]+=32 arr:", arr)
	fmt.Println("after s1[0]+=32 s1:", s1)
	fmt.Println("after s1[0]+=32 s2:", s2)
	s2[1] += 32
	fmt.Println("after s2[1]+=32 arr:", arr)
	fmt.Println("after s2[1]+=32 s1:", s1)
	fmt.Println("after s2[1]+=32 s2:", s2)
	arr[3] += 99
	fmt.Println("after arr[3]+=99 arr:", arr)
	fmt.Println("after arr[3]+=99 s1:", s1)
	fmt.Println("after arr[3]+=99 s2:", s2)
}

//make 创建切片
func MakeSlice() {
	//5:切片的长度 ，初始化之后，没有append的情况，长度之外的空间是不能访问的。
	//6:切片的容量，指的是最大的容量，超过了就翻倍扩容。
	s := make([]int, 5, 6)
	fmt.Printf("make([]int,5,6):%v,addr:%p,len:%d,cap:%d \n", s, s, len(s), cap(s))

	s = append(s, 21)
	fmt.Printf("s=append(s,21):%v,addr:%p,len:%d,cap:%d \n", s, s, len(s), cap(s))

	s = append(s, 21, 32, 43)
	//扩容后，内存地址也改变了。把旧的copy过去
	fmt.Printf("s=append(s,21,32,43):%v,addr:%p,len:%d,cap:%d \n", s, s, len(s), cap(s))

}

func MakeSlice1() {
	s := make([]int, 5, 6)
	//s[5]=100 //数组越界
	s[2] = 99
	fmt.Printf("s=%v \n", s)
	//定义一个空切片，通过append 可以确保切片内的值都是有效的，没有默认的0值了。
	s1 := make([]int, 0, 10)
	fmt.Printf("s1=%v len:%d,cap:%d \n", s1, len(s1), cap(s1))
	//s1[0]=100
	s1 = append(s1, 323)
	fmt.Printf("s1=%v len:%d,cap:%d \n", s1, len(s1), cap(s1))

}

//切片在切片
func SliceExample1() {
	ss := [...]string{"a", "aw", "sa", "asew", "aew", "awe", "aew", "arw", "awr", "aerw"}
	sslice := ss[1:5]
	fmt.Printf("1.sslice=%v len:%d,cap:%d \n", sslice, len(sslice), cap(sslice))
	//切片在切片，只能在前排的起始位置往后继续切片。类似追加，但是切片之前的数据不会在包括，除非在新的切片对原数组
	sslice = sslice[:cap(sslice)]
	fmt.Printf("2.sslice=%v len:%d,cap:%d \n", sslice, len(sslice), cap(sslice))
}

//空切片的append
func SliceExample2() {
	var ss []int
	if ss == nil {
		fmt.Printf("ss is nil ! len:%d,cap:%d !\n", len(ss), cap(ss))
		//空切片不能直接访问，但是可以append
		//长度是正常增长，但是容量是翻倍 ，1，2，4，8
		ss = append(ss, 212)
		fmt.Printf("ss=%v， len:%d,cap:%d !\n", ss, len(ss), cap(ss))
		ss = append(ss, 212)
		fmt.Printf("ss=%v， len:%d,cap:%d !\n", ss, len(ss), cap(ss))
		ss = append(ss, 212)
		fmt.Printf("ss=%v， len:%d,cap:%d !\n", ss, len(ss), cap(ss))

	}
}

//切片的copy,copy 的是交集，谁少以谁的为准。
func SliceCopy() {
	a := []int{3, 4, 5, 6, 7}
	b := []int{44, 55, 445, 545}

	fmt.Printf("befor a=%v \n", a)
	fmt.Printf("befor b=%v \n", b)

	copy(a, b)
	fmt.Printf("after a=%v \n", a)
	fmt.Printf("after b=%v \n", b)
}

//切片汇总
func SumArray(arr []int) int {
	var res int = 0
	for _, v := range arr {
		res += v
	}
	return res
}

//练习
func SliceExample3() {
	//默认有5个空字符串,append 也不会覆盖，只能自己覆盖
	var ss []string = make([]string, 4, 10)
	for i := 0; i < 10; i++ {
		ss = append(ss, fmt.Sprintf("i=%d", i))
	}
	fmt.Printf("ss=%v \n", ss)

}
func main() {

	slice()

	SliceInit()

	//SliceExample()

	SliceExample2()

	arr := [...]int{12, 14, 25, 3, 6, 7, 5, 85, 6, 8}
	fmt.Printf("arr=%v,sum:%d \n", arr, SumArray(arr[:]))

	SliceCopy()

	SliceExample3()
}
