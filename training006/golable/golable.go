package golable

import "fmt"

//break 标签
func ForBreak() {
	fmt.Println("ForBreak")
LOOP1:
	for i := 0; i < 2; i++ {
		fmt.Printf("i:%d \n", i)
		if i == 1 {
			break LOOP1 //没有指定标签，默认跳出循环
		}
	}

	fmt.Println("beak LOOP1")
}

func ForSwithBreak() {
	fmt.Println("ForSwithBreak")
	for {
		x := 1
		switch {
		case x == 3:
			fmt.Printf("x=%d\n", x)
			break // 默认只会跳出一层，
		case x > 0:
			x += 1
		default:
			fmt.Printf("default %d \n", x)

		}
		break // 默认只会跳出一层，
	}
}

func ForSwithContinue() {
	fmt.Println("ForSwithContinue")
LOOP1:
	for i := 0; i < 5; i++ {
		switch {
		case i == 3:
			fmt.Printf("B %d \n", i)
			break LOOP1
		case i < 3:
			fmt.Printf("A %d \n", i)
			continue LOOP1
		default:
			fmt.Printf("default %d \n", i)

		}
		fmt.Printf("i is %d \n", i)
	}

	fmt.Printf("beak LOOP1 \n")
}

func GotoTag() {
	fmt.Println("GotoTag")
	x := 1
	for {
		switch {
		case x == 3:
			fmt.Printf("goto x=%d \n", x)
			goto BREAK
		case x > 0:
			x += 1
		default:
			fmt.Printf("default %d \n", x)

		}
	}
BREAK:
	fmt.Println("go to break ")
}
