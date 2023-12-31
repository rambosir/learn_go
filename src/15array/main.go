package main

import "fmt"

// 数组
// 存放元素的容器
// 存放指定存放元素的类型和容量（长度）
// 数组长度是数组类型的一部分
func main() {

	var a1 [3]bool // [true false true]
	var a2 [4]bool // [true true false false]
	fmt.Printf("a1:%T a2:%T\n", a1, a2)

	// 数组初始化
	// 如果不初始化：默认元素都是零值
	fmt.Println(a1, a2)
	// 1.初始化方式1
	a1 = [3]bool{true, true, true}
	fmt.Println(a1)
	// 2.初始化方式2：根据初始值自动推断数组的长度是多少
	//a100 := [100]int{0, 1}
	a10 := [...]int{0, 1, 2, 3, 4}
	fmt.Println(a10)
	// 3.初始化方式3，根据索引初始化
	a3 := [5]int{0: 1, 4: 2}
	fmt.Println(a3)

	// 数组遍历
	citys := [...]string{"bj", "sh", "sz"}
	for i := 0; i < len(citys); i++ {
		fmt.Println(citys[i])
	}
	for i, city := range citys {
		fmt.Println(i, city)
	}

	// 多维数组
	//[[12] [2 3] [4 5]]
	var a11 [3][2]int8
	a11 = [3][2]int8{
		[2]int8{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a11)
	//多为数组的遍历
	for _, v := range a11 {
		fmt.Println(v)
		for _, v2 := range v {
			fmt.Println(v2)
		}

	}
	// 数组是值类型
	b1 := [...]int{1, 2, 3}
	b2 := b1 // ctrl+c ctrl+v => windows复制粘贴
	b2[0] = 100
	fmt.Println(b1, b2)
}
