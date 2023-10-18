package main

import "fmt"

// for 循环
func main() {

	// 基本格式
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("-----------")
	//变种1
	var i = 5
	for ; i < 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("-----------")
	//变种2
	var j = 5
	for j < 10 {
		fmt.Println(j)
		j++
	}

	// 无限循环
	//for {
	//	fmt.Println("123")
	//}

	// for range循环
	s := "hello world世界"
	for k, v := range s {
		fmt.Printf("%d %c\n", k, v)
	}
}
