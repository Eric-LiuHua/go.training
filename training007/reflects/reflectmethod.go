package reflects

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"training007/interfaces"
)

//to string 接口
type Stringer interface {
	ToString() string
}

//温度
type Celsius float64

func (c Celsius) ToString() string {
	return strconv.FormatFloat(float64(c), 'f', 1, 64) + " °C"
}

type Day int

//切片不能做常量
var dayname = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

func (d Day) ToString() string {
	return dayname[((int)(reflect.ValueOf(d).Int()))%len(dayname)]
}

func MyPrint(args ...interface{}) {
	for i, v := range args {
		if i > 0 {
			os.Stdout.WriteString(" ")
		}
		switch t := v.(type) {
		case string:
			os.Stdout.WriteString(t)
		case Stringer:
			os.Stdout.WriteString(t.ToString())
		case int:
			os.Stdout.WriteString(strconv.Itoa(t)) //转换
		default:
			os.Stdout.WriteString("???")
		}
	}
	os.Stdout.WriteString(" \n")
}

//需要调用函数，需要value
func StructInfoValue(o interface{}) {

	fmt.Println("StructInfoType reflect.TypeOf(o):", reflect.TypeOf(o))

	if reflect.TypeOf(o).Kind() == reflect.Struct {
		funcValue := reflect.ValueOf(o)
		funcType := reflect.TypeOf(o)

		fmt.Println("funcValue.NumMethod() at Struct:", funcValue.NumMethod())
		for i := 0; i < funcValue.NumField(); i++ {
			os.Stdout.WriteString(fmt.Sprintf("field [%d] %s ,type:%T ,tag:%v , value:%v \n", i, funcType.Field(i).Name, funcValue.Field(i), funcType.Field(i).Tag, funcValue.Field(i)))
		}
		for j := 0; j < funcValue.NumMethod(); j++ {
			os.Stdout.WriteString(fmt.Sprintf("Method at value [%d] %s ,type:%T  \n", j, funcType.Method(j).Name, funcValue.Method(j)))
			funcValue.Method(j).Call([]reflect.Value{})
		}
		if funcValue.MethodByName("FuncName1").IsValid() {
			funcValue.MethodByName("FuncName1").Call([]reflect.Value{})
		}
	} else if reflect.TypeOf(o).Kind() == reflect.Ptr {
		//方法如果在指针上，那就需要Valueof(指针) ,在值上就要valuof(值)
		funcValue := reflect.ValueOf((o))
		fmt.Println("funcValue.NumMethod() at ptr:", funcValue.NumMethod())

		for j := 0; j < funcValue.NumMethod(); j++ {
			os.Stdout.WriteString(fmt.Sprintf("Method at ptr [%d] ,type:%T value:%v \n", j, funcValue.Method(j), reflect.TypeOf(funcValue.Method(j)).Name()))
			res := funcValue.Method(j).Call([]reflect.Value{})
			if len(res) > 0 {
				fmt.Println("funcValue.Method(j).Call([]reflect.Value{}):", res[0])
			}
		}
		if funcValue.MethodByName("FuncName1").IsValid() {
			funcValue.MethodByName("FuncName1").Call([]reflect.Value{})
		}

	} else {
		os.Stdout.WriteString(fmt.Sprintf("arg is not struct ! type:%v \n", reflect.TypeOf(o).Kind()))
	}
}

//获取字段和函数信息，用typeof
func StructInfoType(o interface{}) {

	fmt.Println("StructInfoType reflect.TypeOf(o):", reflect.TypeOf(o))

	if reflect.TypeOf(o).Kind() == reflect.Struct {
		funcType := reflect.TypeOf(o)
		fmt.Println("funcType.NumMethod() at Struct:", funcType.NumMethod())
		for i := 0; i < funcType.NumField(); i++ {
			os.Stdout.WriteString(fmt.Sprintf("field [%d] ,type:%T ,tag:%v  value:%v \n", i, funcType.Field(i), funcType.Field(i).Tag, funcType.Field(i).Type))
		}
		for j := 0; j < funcType.NumMethod(); j++ {
			os.Stdout.WriteString(fmt.Sprintf("Method at value [%d],name:%v ,type:%T value:%v \n", j, funcType.Method(j).Name, funcType.Method(j), funcType.Method(j).Type))
		}

	} else if reflect.TypeOf(o).Kind() == reflect.Ptr {
		//方法如果在指针上，那就需要Valueof(指针) ,在值上就要valuof(值)
		funcType := reflect.TypeOf((o))
		fmt.Println("funcValue.NumMethod() at ptr:", funcType.NumMethod())

		for j := 0; j < funcType.NumMethod(); j++ {
			os.Stdout.WriteString(fmt.Sprintf("Method at value [%d],name:%v ,type:%T value:%v \n", j, funcType.Method(j).Name, funcType.Method(j), funcType.Method(j).Type))
		}

	} else {
		os.Stdout.WriteString(fmt.Sprintf("arg is not struct ! type:%v \n", reflect.TypeOf(o).Kind()))
	}
}

func TestMyPrint() {
	MyPrint(Day(21), "was", Celsius(18.36))
	//指针的kind 是ptr
	StructInfoType(&interfaces.Rectangle{Width: 18, Lenght: 23})
	StructInfoType(RectangleA{width: 18, lenght: 23})

	StructInfoValue(&interfaces.Rectangle{Width: 18, Lenght: 23})
	StructInfoValue(RectangleA{width: 18, lenght: 23})
}
