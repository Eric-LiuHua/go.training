package arrayex

import (
	"fmt"
)

//数组到长度也是数组的类型。不同类型的不能

//默认满列表，默认下标 ，
func testArray() {
	var a [3]int = [3]int{1, 2, 3}
	fmt.Println(a)

}

//通过初始化确定长度
func testArrayb() {
	fArr := [...]int{1, 323, 22}
	fmt.Println(fArr)
}

//通过下标指定
func testArraya() {
	var a [5]int = [5]int{0: 1, 3: 2, 4: 3}
	for i := 0; i < len(a); i++ {
		println(a[i])
	}
}

//二维数组
func eTestAyyay() {
	var earr [3][2]int
	earr[0][0] = 1
	earr[0][1] = 2
	earr[1][0] = 3
	earr[1][1] = 4
	earr[2][0] = 5
	earr[2][1] = 6

	//第二种声明方式
	var earr1 [3][2]int = [3][2]int{{123, 434}, {345, 234}, {23, 24}}

	for i := 0; i < len(earr1); i++ {
		for j := 0; j < len(earr1[0]); j++ {
			fmt.Printf("earr1[%d][%d]=%d \n", i, j, earr1[i][j])
		}
	}

}

//...的使用,多个可以接受任意个int参数
func testddd(a ...int) {
	//切片声明
	var b = []int{
		1,
		2,
		32,
	}
	for _, x := range a {
		fmt.Printf("dddd:%d \n", x)
	}
	b = append(b, a...) //a的元素被打散一个个append进b
	fmt.Println(b)

}

//切片的...使用
func testdddi() {

	var strss = []int{
		1,
		2,
		32,
	}
	var strss2 = []int{
		32,
		43,
		3213,
		231,
	}
	strss = append(strss, strss2...) //strss2的元素被打散一个个append进strss
	fmt.Println(strss)

}

//求和,没有5 就是切片了
//数组的缺陷，参数都是必须带长度带，而且不同长度带不能想java 那样
func sumArray(arr [5]int) int {
	var res int = 0
	for _, v := range arr {
		res += v
	}
	return res
}

//两数子和
func twoSum(arr []int, target int) {
	pass := make(map[int]int)
	for i, v := range arr {
		tmp := target - v
		tv, ok := pass[tmp]
		//fmt.Println("tv:", tv)
		//fmt.Println("ok:", ok)
		//fmt.Println("tmp:", tmp)
		if ok {
			fmt.Printf("res(%d,%d)\n", tv, i)
		}
		pass[v] = i
	}
}

func ArrayMain() {
	testArraya()
	testArrayb()
	eTestAyyay()

	testddd(1, 2, 3, 2, 3)
	testdddi()

	var a [5]int = [5]int{0: 1, 3: 2, 4: 3}

	b := [...]int{1, 3, 5, 8, 7}

	fmt.Println("sumArray(a):", sumArray(a))
	fmt.Println("twoSum(b[:], 8):", b)
	twoSum(b[:], 8)
}
