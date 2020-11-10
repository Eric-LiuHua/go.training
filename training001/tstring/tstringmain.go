package tstring

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

/**
*  字符串底层就是byte 数组，可以和[]byte 切片相互转换。
*  字符串不能修改值，要修改可以转为切片，修改结束在转会字符。其他就是另外的字符串来
*
*/
func TestString(){
	fmt.Println("............... tstring.TestString  start ...............")

	str := "字符串底层就是byte 数组，可以和[]byte 切片相互转换。"
	//
	str2 :=[]rune(str)
	fmt.Printf("不正常 str[0]=%c,len=%d \n",str[0],len(str))
	fmt.Printf("正常 str2[0]=%c,len=%d \n",str2[0],len(str2))


	fmt.Println("............... for i,v :=  range str  中文正常...............")

	for i,v :=  range str {
		fmt.Printf("str[%d]=%c\n",i,v)
	}
	fmt.Println("............... for i := 0; i < len(str2); i++  str2 :=[]rune(str) 转化为切片后中文正常  ...............")

	for i := 0; i < len(str2); i++ {
		fmt.Printf("str2[%d]=%c\n",i,str2[i])
	}


	var strSlice= []byte(str)
	strSlice[0]='x'
	str=string(strSlice)
	fmt.Printf("str=%s\n",str)


	fmt.Println("............... tstring.TestString  end ...............")

}

//golang中string底层是通过byte数组实现的。
//中文字符在unicode下占2个字节，在utf-8编码下占3个字节，而golang默认编码正好是utf-8。

//golang中海油一个byte数据类型与rune相似，它们都是用来表示字符类型的变量类型。它们的不同在于：
//byte 等同于int8，常用来处理ascii字符
//rune 等同于int32,常用来处理unicode或utf-8字符
func TestRune()  {
	fmt.Println("............... TestRune  start ...............")
	var str = "hello 你好！"
	 //golang中string底层是通过byte数组实现的，直接求len 实际是在按字节长度计算  所以一个汉字占3个字节算了3个长度
	 fmt.Println("期望是9，5个字母，一个空格，两个中文，一个 感叹号 ,len(str):", len(str))

	 fmt.Println("golang中的unicode/utf8包提供了用utf-8获取长度的方法,utf8.RuneCountInString(str):", utf8.RuneCountInString(str))
 
	 fmt.Println("通过rune类型处理unicode字符,len([]rune(str)):", len([]rune(str)))
	 fmt.Println("............... TestRune  end ...............")
}

//golang字符串比较的三种常见方法
func TestEqual()  {
	fmt.Println("............... TestEqual  start ...............")
	
	fmt.Println("1.自建方法“==”，区分大小写，最简单的方法:")

	fmt.Println("\"go\"==\"go\"","go"=="go")
	fmt.Println("\"GO\"==\"go\"","GO"=="go")
	
	fmt.Println("2.Compare函数，区分大小写，比自建方法“==”的速度要快，string comparison operators ==, <, >, and so on:")

	fmt.Println("strings.Compare(\"GO\",\"go\")",strings.Compare("GO","go"))
	fmt.Println("strings.Compare(\"go\",\"go\")",strings.Compare("go","go"))


	fmt.Println("3.比较UTF-8编码在小写的条件下是否相等，不区分大小写，are equal under Unicode case-folding.:")
	fmt.Println("strings.EqualFold(\"GO\",\"go\"):",strings.EqualFold("GO","go"))

	fmt.Println("............... TestEqual  start ...............")
}


//倒序输出
func RreverseString(str string) string  {
 
	var srune =[]rune(str)
	slen:=len(srune)
	if slen<=1{
		return str
	}else{
		for i,j:= 0,slen-1; i < j; i,j=i+1 ,j-1 {
			 srune[i],srune[j]=srune[j],srune[i]
		}
	}
	return string(srune)
}

//是否回文数
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
