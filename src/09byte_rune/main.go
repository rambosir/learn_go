package main

import "fmt"

// byte 和rune类型
// go语言中为了处理非ASCII码类型的字符 定义了新的rune类型
func main() {
	s := "hello world"

	n := len(s)
	fmt.Println(n)

	//for i := 0; i < len(s); i++ {
	//fmt.Println(s[i])
	//	fmt.Printf("%c\n1", s[i]) // %c:字符
	//}

	//for _, c := range s {
	//	fmt.Printf("%c\n1", c)
	//}

	// 字符串修改
	s2 := "白萝卜"      // => '白' '萝' '卜'
	s3 := []rune(s2) // 把字符串强制转换成rune切片
	s3[0] = '红'
	fmt.Println(string(s3))

	c1 := "红"
	c2 := '红' // rune(int32)
	fmt.Printf("c1:%T c2:%T\n", c1, c2)
	c3 := "h"
	c4 := 'h'
	fmt.Printf("c3:%T c4:%T\n", c3, c4)
	fmt.Printf("c4:%d\n", c4)

	// 类型转换
	n1 := 10
	var f float64
	f = float64(n1)
	fmt.Println(f)
	fmt.Printf("%T\n", f)
}
