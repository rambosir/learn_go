package main

import "fmt"

//  go语言推荐使用驼峰式命名

// var student_name string
var studentName string

//var StudentName string

// 声明变量
//var name string
//var age int
//var isOk bool

// 批量声明变量
var (
	name string // ""
	age  int    // 0
	isOk bool   // false
)

func main() {
	name = "理想"
	age = 18
	isOk = true
	// go语言中变量声明必须使用，不使用编译不通过
	fmt.Print(isOk) // 不换行打印
	fmt.Println()
	fmt.Printf("name:%s\n", name) // %s:占位符
	fmt.Println(age)              // 换行打印

	// 声明变量同时赋值
	var s1 string = "王XX"
	fmt.Println(s1)
	// 类型推导（根据值判断该变量类型）
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明，只能在函数里面使用
	s3 := "haha"
	fmt.Println(s3)
	// s1:="10" 同一作用域不能重复声明同名的变量
	// 匿名变量

	_, y := test()
	fmt.Println(y)
}

func test() (int, string) {
	return 0, ""
}
