package main

import "fmt"

func main() {
	b1 := true
	var b2 bool // 默认false
	fmt.Println("%T\n", b1)
	fmt.Println("%T value:%v\n", b2, b2)
}
