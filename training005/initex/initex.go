package initex

import (
	_ "ex001" //_ 表示要引用这个包，但是我不需要调用这个包的内容
	"fmt"
)

var g int = 10

func init() {
	fmt.Println("initex start init ,g=", g)
	g = 11
}
func Tmp() {
	fmt.Println("main start ,g=", g)
}

func main() {
	Tmp()

}
