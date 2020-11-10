package reflects

import (
	"fmt"
	"reflect"
)

type Object struct {
	Id   int
	Code int
	Name string
}

//.Int 和 String 可以帮助我们分别取出 reflect.Value 为 int64 和 string的值
func TestReflectBase() {
	fmt.Println("TestReflectBase start ...............")

	x := "3.14"
	fmt.Printf("		reflect.TypeOf(x):%v \n", reflect.TypeOf(x))
	fmt.Printf("		reflect.ValueOf(x):%v \n", reflect.ValueOf(x))
	fmt.Printf("		reflect.ValueOf(x).String():%v \n", reflect.ValueOf(x).String())
	fmt.Printf("		reflect.TypeOf(x).Kind:%v \n", reflect.TypeOf(x).Kind())
	fmt.Printf("		reflect.ValueOf(x).Kind:%v \n", reflect.ValueOf(x).Kind())

	fmt.Printf("		reflect.ValueOf(x).CanSet:%v \n", reflect.ValueOf(x).CanSet())
	fmt.Printf("		reflect.ValueOf(x).CanAddr:%v \n", reflect.ValueOf(x).CanAddr())
	fmt.Printf("		reflect.ValueOf(x).CanInterface:%v \n", reflect.ValueOf(x).CanInterface())

}

//elem 赋值使用
func TestReflectElem() {
	x := "3.14"
	vaddr := reflect.ValueOf(&x)
	vm := vaddr.Elem()

	fmt.Println("TestReflectElem start ...............")
	fmt.Printf("take the address of reflect.ValueOf(x):%v \n", reflect.ValueOf(x))
	fmt.Printf("take the address of reflect.ValueOf(&x):%v \n", vaddr)
	fmt.Printf("take the address of &x:%v \n", &x)
	fmt.Printf("take the address of reflect.ValueOf(x).CanAddr:%v \n", reflect.ValueOf(x).CanAddr())
	vm.SetString("3.1415") // this works!
	fmt.Println("vm.Interface():", vm.Interface())
	fmt.Println("vm:", vm)

}

func TestReflectStruct() {
	o := Object{
		Id:   1,
		Code: 102,
		Name: "strdsfs",
	}
	PrintTypeKind(o)
	PrintNamefields(o)
}

func PrintTypeKind(p interface{}) {
	fmt.Println("	rflect.TypeOf(p) ", reflect.TypeOf(p))
	fmt.Println("	reflect.TypeOf(p).Kind ", reflect.TypeOf(p).Kind())
	fmt.Println("	reflect.ValueOf(p).Kind ", reflect.ValueOf(p).Kind())
	fmt.Println("	reflect.TypeOf(p).Name ", reflect.TypeOf(p).Name())
}

func PrintNamefields(p interface{}) {
	if reflect.TypeOf(p).Kind() == reflect.Struct {
		fields := reflect.ValueOf(p)
		fmt.Println("Number of fields", fields.NumField())
		for i := 0; i < fields.NumField(); i++ {
			fmt.Printf("fields[%d] type:%T value:%v \n", i, fields.Field(i), fields.Field(i))
		}
	} else {
		fmt.Println("parameter is not reflect.Struct,type =", reflect.TypeOf(p).Kind())
	}
}

//所有结构体都可以转化为空接口 func TypeOf(i interface{}) Type
func TestReflect() {
	TestReflectBase()
	TestReflectStruct()
}
