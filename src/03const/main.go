package main

import "fmt"

// 常量
// 定义了常量后不能再修改
const pi = 3.1415926

// 批量声明常量
const (
	statusOk = 200
	notFound = 404
)

// 批量声明，如果某一行声明后没有赋值，默认就和上一行一致
const (
	n1 = 100
	n2
	n3
)

// iota：类似枚举
const (
	a1 = iota //
	a2
	a3
)

const (
	b1 = iota
	b2
	_
	b3
)

// 插队
const (
	c1 = iota
	c2 = 100
	c3 = iota
	c4
)

// 多个常量声明在一行
const (
	d1, d2 = iota + 1, iota + 2
	d3, d4 = iota + 1, iota + 2
)

// 定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

const (
	a, b = iota + 1, iota + 2 //1,2
	c, d                      //2,3
	e, f                      //3,4
)

func main() {
	fmt.Println(n1, n2, n3)

	fmt.Println(a1, a2, a3)

	fmt.Println(b1, b2, b3)

	fmt.Println(c1, c2, c3, c4)

	fmt.Println(d1, d2, d3, d4)

	fmt.Println(a, b, c, d, e, f)
}
