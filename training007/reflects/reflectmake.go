package reflects

import (
	"fmt"
	"reflect"
)

func TestMakeFun() {
	swap := func(in []reflect.Value) []reflect.Value {
		return []reflect.Value{in[1], in[0]}
	}
	makeSwap := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		v := reflect.MakeFunc(fn.Type(), swap)
		fn.Set(v)
	}

	var intSwap func(int, int) (int, int)
	makeSwap(&intSwap)
	//这个就是将参数0,1传入，运行swap
	fmt.Println(intSwap(0, 1)) //返回：1 0

	var floatSwap func(float64, float64) (float64, float64)
	makeSwap(&floatSwap)
	fmt.Println(floatSwap(2.72, 3.14)) //返回：3.14 2.72

}

//使用new来创建map时，返回的内容是一个指针，这个指针指向了一个所有字段全为0的值map对象，
//需要初始化后才能使用，而使用make来创建map时，返回的内容是一个引用，可以直接使用
func MakeMap() {
	imap := make(map[string]string)
	var input interface{}
	input = imap
	m := reflect.ValueOf(input)
	if m.Kind() == reflect.Map {
		res := reflect.MakeMap(m.Type())
		keys := m.MapKeys()
		for _, k := range keys {
			key := k.Convert(res.Type().Key()) //.Convert(m.Type().Key())
			value := m.MapIndex(key)
			res.SetMapIndex(key, value)
		}
	}
}
