package main

import "fmt"

// 浮点型

func main() {

	//math.MaxFloat32 // float32最大值
	var f1 = 1.23456
	fmt.Printf("%T\n", f1) // 默认go语言中小数是float64类型
	f2 := float32(1.23456)
	fmt.Printf("%T\n", f2) // 显示声明float32类型

	//f1 = f2 float32类型不能直接赋值给float64类型变量
}
