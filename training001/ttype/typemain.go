package ttype

import (
	"fmt"
	"strings"
)

func TestInt8() {
	fmt.Println("............... TestInt8 start ...............")
	fmt.Println("TestInt8")
	var a int8
	a = 18
	fmt.Println("a=", a)
	a = -12
	fmt.Println("a=", a)

	var b int
	b = 121423123
	fmt.Println("b= ", b)

	//go 强类型，类型不一致的需要转换
	b = int(a)
	fmt.Println(" b=int(a) = ", b)

	fmt.Println("== 操作符 b==a  ", b == int(a))
	b = b + 1
	fmt.Println("b=b+1,b= ", b)

	fmt.Println("............... TestInt8  end ...............")
}

func TestString() {
	fmt.Println("............... TestString  start ...............")
	var a string
	a = "a:hello go "
	fmt.Println("a= ", a)
	b := "b:hello go "
	fmt.Printf("b= %v \n", b)
	c := `c:hello go
	dslksjf\n
	dsfs 
	使用反引号来定义多行字符串  ,英文状态的 ～ 
	go run  *.go 不需要build
	./* 执行可执行文件，需要go build  `
	fmt.Printf("c= %v \n", c)

	clen := len(c)
	fmt.Printf("c.len= %v \n", clen)

	fmt.Println("............... TestString  end ...............")

}

//切片测试
func TestSlice() {
	fmt.Println("............... TestSlice  start ...............")

	var a string
	a = "a:hello go "
	b := "b:hello go "

	// Sprintf 字符串组合
	s := fmt.Sprintf("Sprintf:%s,b=%s", a, b)

	fmt.Printf("s= %v \n", s)

	// 字符串分割
	ids := "12,23,1,42,4,34,54,35,34,53,534,99"

	fmt.Println(`
	数组需要明确指定大小，切片不需要。数组是值传递，切片是地址传递。
	数组： a := [3]int{1,2,3}   或  var a  = [3]int {1, 2, 3}

	切片： a:= []int{1,2,3}    或 var a = []int{1,2,3}
	a := make([]int, 5)  

 	a := make([]int, 5, 10)    

	值得注意的是： string 类型 的底层是数组， 是数组就可以用来切片操作
	`)
	id_slice := strings.Split(ids, ",")
	fmt.Println("result:", id_slice)
	fmt.Println("len:", len(id_slice))
	fmt.Println("id_slice[0,1]:", id_slice[0:1])
	fmt.Println("cap:", cap(id_slice))

	//长度为5，容量为10
	//如果新的长度小于容量，不会更换底层数组,否则拷贝过去，原数组丢弃
	//容量的用途是：在数据拷贝和内存申请的消耗与内存占用之间提供一个权衡。
	a_slice := make([]int, 5, 10)
	fmt.Println("a_slice:", a_slice)

	//长度为5，容量为5
	b_slice := make([]int, 5)

	fmt.Println("b_slice:", b_slice)

	fmt.Println("............... TestSlice  end ...............")

}

//strings 普通检查校验
func TestStrings() {
	fmt.Println("............... TestStrings  start ...............")

	url := "https://www.baidu.com"

	fmt.Printf("是否包含 strings.Contains(%s,\"http\") ,result:%v \n", url, (strings.Contains(url, "http:")))

	fmt.Printf("前缀检查 strings.HasPrefix(%s,\"http\") ,result:%v \n", url, (strings.HasPrefix(url, "http:")))
	fmt.Printf("前缀检查 strings.HasPrefix(%s,\"https\") ,result:%v \n", url, (strings.HasPrefix(url, "https:")))

	fmt.Printf("后缀检查 strings.HasSuffix(%s,\"com\") ,result:%v \n", url, (strings.HasSuffix(url, ".com")))
	fmt.Printf("后缀检查 strings.HasSuffix(%s,\"org\") ,result:%v \n", url, (strings.HasSuffix(url, ".org")))

	fmt.Printf("Index 检查 strings.Index(%s,\"com\") ,result:%v \n", url, (strings.Index(url, ".com")))
	fmt.Printf("LastIndex 检查 strings.Index(%s,\"com\") ,result:%v \n", url, (strings.LastIndex(url, ".com")))

	//切片声明
	var urls = []string{"git://github.com/rogpeppe/godef.git", "git://github.com/kisielk/errcheck.git", "https://www.github.com"}

	urlc := strings.Join(urls, ";")
	fmt.Println("字符合并 urlc:", urlc)
	fmt.Println("............... TestStrings  end ...............")

}

func TestFloat64() {

	fmt.Println("******* TestFloat64 ************ ")
	var v = 67.6
	fmt.Println(int64(v * 100))
	fmt.Println(int64(v * 10 * 10))

}
