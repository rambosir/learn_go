package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业啦！")
	} else {
		fmt.Println("该写暑假作业啦！")
	}

	// 多个判断条件
	if age > 35 {
		fmt.Println("人到中年")
	} else if age > 18 {
		fmt.Println("青年")
	} else {
		fmt.Println("少年")
	}
}

func ifDemo2() {
	// 作用域
	// age变量此时只在if条件判断语句中生效
	if score := 65; score >= 90 {
		fmt.Println("A")
	} else if score > 75 {
		fmt.Println("B")
	} else {
		fmt.Println("C")
	}

}
