package main

import (
	"fmt"
	"strings"
)

// 字符串
func main() {

	path := "\"D:\\\""
	fmt.Println(path)

	s := "i'm ok"
	fmt.Println(s)

	//多行字符串
	s2 := `
这是一个多行字符串。
换行
`
	fmt.Println(s2)
	s3 := `D:\a\b\c`
	fmt.Println(s3)

	// 字符串相关操作
	fmt.Println(len(s3))
	name := "hello"
	world := "world"
	ss := name + world
	fmt.Println(ss)
	ss1 := fmt.Sprintf("%s%s", name, world)
	fmt.Println(ss1)
	// fmt.Printf("%s%s", name, world)
	// 分割
	ret := strings.Split(s3, "\\")
	fmt.Println(ret)

	// 包含
	fmt.Println(strings.Contains(ss, "hello1"))
	fmt.Println(strings.Contains(ss, "hello"))
	// 前缀
	fmt.Println(strings.HasPrefix(ss, "hello"))
	// 后缀
	fmt.Println(strings.HasSuffix(ss, "hello"))

	s4 := "abcdeb"
	fmt.Println(strings.Index(s4, "c"))
	fmt.Println(strings.LastIndex(s4, "b"))

	// 拼接
	fmt.Println(strings.Join(ret, "+"))

}
