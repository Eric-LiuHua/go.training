package mapex

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

// map 是引用类型
func MapTest() {
	countryCapitalMap := make(map[string]string)
	/* map插入key - value对,各个国家对应的首都 */
	countryCapitalMap["France"] = "巴黎"
	countryCapitalMap["Italy"] = "罗马"
	countryCapitalMap["Japan"] = "东京"
	countryCapitalMap["India"] = "新德里"
	countryCapitalMap["China"] = "北京"

	for c := range countryCapitalMap {
		fmt.Printf("%s 首都 ：%s \n", c, countryCapitalMap[c])
	}
	/*删除元素*/
	delete(countryCapitalMap, "France")
	fmt.Println("法国条目被删除")

	fmt.Println("删除元素后地图")
	for c := range countryCapitalMap {
		fmt.Printf("%s 首都 ：%s \n", c, countryCapitalMap[c])
	}

	// 判断map 对key 是否存在
	v1, isexist1 := countryCapitalMap["cccc"]
	fmt.Printf("v1:%v,isexist1:%v \n", v1, isexist1)
	v2, isexist2 := countryCapitalMap["China"]
	fmt.Printf("v2:%v,isexist2:%v \n", v2, isexist2)
}

//涉及排序,使用切片进行排序
func MapTest1() {
	rand.Seed(time.Now().UnixNano())
	mapa := make(map[string]int)
	var slice []string = make([]string, 0, 128)
	n := 1000
	for i := 0; i < 128; i++ {
		mapa[fmt.Sprintf("key:%d", i)] = rand.Intn(n)
	}
	fmt.Println("************** mapa  *********************")
	for mindex, mvalue := range mapa {
		fmt.Printf("m:%s = %d \n", mindex, mvalue)
		slice = append(slice, mindex)
	}

	sort.Strings(slice)

	fmt.Println("************** sort.Strings(slice) *********************")
	for _, v := range slice {
		fmt.Printf("k:%s ,value:%d \n", v, mapa[v])
	}
}

//, cap(mapa) 不能使用，
func MapTest2() {
	//mapa 是引用类型，本身就是地址。不需要传地址
	mapa := make(map[int]int, 22)
	fmt.Printf("1.mapa=%v ,len=%d \n", mapa, len(mapa))

	mapa[1] = 23
	mapa[3] = 23
	mapa[14] = 23

	mapb := mapa

	//同时更新mapa的
	mapb[1] = 2332
	mapb[3] = 23324
	fmt.Printf("2.mapa=%v ,len=%d \n", mapa, len(mapa))
	fmt.Printf("2.mapb=%v ,len=%d \n", mapb, len(mapb))
}

//map 类型的切片使用
func MapTest3() {

	//切片没有初始化长度，那就需要append 来追加数据，进行默认扩张。
	var slice []map[string]int = make([]map[string]int, 5, 128)
	//使用map的时候，还需要make 初始化

	slice[0] = make(map[string]int, 11)

	slice[0]["aa"] = 99
	slice[0]["c"] = 99
	slice[0]["cc"] = 99
	slice[0]["sf"] = 99

	for sindex, smpa := range slice {
		fmt.Printf("sindex:%d,smap:%v \n", sindex, smpa)
	}
}

//切片 类型的map使用
func MapTest4() {
	rand.Seed(time.Now().UnixNano())
	//map 的值是切片，因为map 没有cap ，所以不需要定义，只需要定义长度即可
	var mapslice map[string][]int = make(map[string][]int, 128)
	//使用map的时候，还需要make 初始化
	var key string = "tmpkey"

	_, isexist := mapslice[key]
	if !isexist {
		mapslice[key] = make([]int, 5, 10)
	}
	mapslice[key] = append(mapslice[key], rand.Intn(900))
	mapslice[key] = append(mapslice[key], rand.Intn(900))
	mapslice[key] = append(mapslice[key], rand.Intn(900))
	for sindex, smpa := range mapslice {
		fmt.Printf("sindex:%s,smap:%v \n", sindex, smpa)
	}
}
