package reflects

import "fmt"

type RectangleA struct {
	lenght float64
	width  float64
}

func (r RectangleA) Area() float64 {
	return r.lenght * r.width
}

func (r RectangleA) FuncName1() {
	fmt.Println("(r RectangleA).FuncName1 被调用")
}
