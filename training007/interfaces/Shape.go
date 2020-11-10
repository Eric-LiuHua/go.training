package interfaces

import "fmt"

type Square struct {
	side float64
}

func (s *Square) Area() float64 {
	return s.side * s.side
}

type Rectangle struct {
	Lenght float64
	Width  float64
}

func (r *Rectangle) Area() float64 {
	return r.Lenght * r.Width
}

//接口类型的参数，这里不能传指针
func GetArea(o Shaper) {
	fmt.Printf("o.Area()：%f \n", o.Area())
	//类型分支
	switch t := o.(type) {
	case *Square:
		fmt.Printf("o type :Square %T ,area:%f  \n", t, o.Area())
	case *Rectangle:
		fmt.Printf("o type :Rectangle %T,area:%f  \n", t, o.Area())
	default:
		fmt.Printf("未知类型! type:%T \n", t)
	}

	//判断类型
	if u, ok := o.(*Rectangle); ok {
		fmt.Printf("o type :Rectangle %T,area:%f  \n", u, o.Area())
	} else {
		fmt.Printf("o  不含类型未Rectangle的变量\n")
	}
}

func ClassIfier(items ...interface{}) {
	for i, v := range items {
		switch t := v.(type) {
		case bool:
			fmt.Printf("参数 #%d 类型是 bool \n", i)
		case float64:
			fmt.Printf("参数 #%d 类型是 float64 \n", i)
		case int64:
			fmt.Printf("参数 #%d 类型是 int64 \n", i)
		case string:
			fmt.Printf("参数 #%d 类型是 string \n", i)
		case nil:
			fmt.Printf("参数 #%d 类型是 nil \n", i)
		default:
			fmt.Printf("参数 #%d 未知类型! type:%T \n", i, t)
		}
	}
}
func ClassIfierI(items ...interface{}) {
	for i, v := range items {
		switch t := v.(type) {
		default:
			fmt.Printf("参数 #%d 类型是 %T \n", i, t)
		}
	}
}

func TestMain() {
	//func new(Type) *Type
	//返回的是指针
	s := new(Square)
	s.side = 100

	r := &Rectangle{Width: 18, Lenght: 23}
	ss := &Square{side: 22}

	shapes := []Shaper{s, ss, r}
	//这样方便知道 是把对象转接口。调用接口的方法
	areaintf := Shaper(s)
	fmt.Printf("areaintf.Area()：%f \n", areaintf.Area())

	GetArea(r)
	GetArea(s)

	for n, _ := range shapes {
		fmt.Printf("形状的参数：%v \n", shapes[n])
		fmt.Printf("形状的面积：%f \n", shapes[n].Area())
	}

	ClassIfier(12, 23.21, "Dfer", nil, s, r, false)
}
