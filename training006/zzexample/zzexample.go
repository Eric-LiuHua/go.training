package zzexample

import "fmt"

type Rect struct {
	//变量名也要首字母大写，否则外部包无法访问
	Width  float64
	Height float64
}

func (r Rect) mj() float64 {
	return r.Width * r.Height
}

func Testmain() {

	//1. &是取地址符号, 取到Rect类型对象的地址
	//2. *可以表示一个变量是指针类型(r是一个指针变量):
	var r *Rect = &Rect{10, 1000}
	fmt.Println("var r *tmp.Rect = &tmp.Rect{10, 1000}")
	fmt.Printf("r:%v\n", r)
	//3.*也可以表示指针类型变量所指向的存储单元 ,也就是这个地址所指向的值
	fmt.Printf("*r:%v\n", *r)
	//4.查看这个指针变量的地址 , 基本数据类型直接打印地址
	fmt.Printf("&r:%v\n", &r)

	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
}
