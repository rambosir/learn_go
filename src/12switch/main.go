package main

import "fmt"

// switch
// 简化大量的if判断
func main() {
	n := 5
	switch n {
	case 1:
		fmt.Println("1")
	case 2:
		fmt.Println("2")
	case 3:
		fmt.Println("3")
	case 4:
		fmt.Println("4")
	case 5:
		fmt.Println("5")
	default:
		fmt.Println("无效输入")

	}

	testSwitch3()
	testSwitch4()
	testSwitch5()
}

func testSwitch3() {
	// 一个分支可以有多个值，多个case值中间使用英文逗号分隔
	switch n := 7; n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func testSwitch4() {
	// 不需要再跟判断变量
	age := 30
	switch {
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

func testSwitch5() {
	// 执行满足条件的case的下一个case，是为了兼容C语言中的case设计的。
	s := "a"
	switch {
	case s == "a":
		fmt.Println("a")
		fallthrough
	case s == "b":
		fmt.Println("b")
	case s == "c":
		fmt.Println("c")
	default:
		fmt.Println("...")
	}
}
