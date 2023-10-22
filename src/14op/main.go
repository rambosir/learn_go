package main

import "fmt"

func main() {

	var (
		a = 5
		b = 2
	)
	fmt.Println(a + b)
	fmt.Println(a - b)
	fmt.Println(a * b)
	fmt.Println(a / b)
	fmt.Println(a % b)
	a++ //单独的语句，不能放在=的右边赋值 ==> a=a+1
	b-- // 单独的语句，不能放在=的右边赋值 ==> b = b-1
	fmt.Println("------")
	// 关系运算符
	fmt.Println(a == b) // go语言是强类型，相同类型的变量才能比较
	fmt.Println(a != b)
	fmt.Println(a >= b)
	fmt.Println(a > b)
	fmt.Println(a <= b)
	fmt.Println(a < b)
	fmt.Println("------")
	//逻辑运算符
	age := 22
	if age > 18 && age < 60 {
		fmt.Println("打工人")
	} else {
		fmt.Println("幸福")
	}
	fmt.Println("------")
	if age < 18 || age > 60 {
		fmt.Println("幸福")
	} else {
		fmt.Println("打工人")
	}

	fmt.Println("------")
	// not取反
	isTure := false
	fmt.Println(isTure)
	fmt.Println(!isTure)
	// 位运算：针对的是二进制
	// 5的二进制：101
	// 2的二进制：10
	// & 按位与 (两位均为1才为1)
	fmt.Println(5 & 2)
	// | 按位或（两位有1个为1就为1）
	fmt.Println(5 | 2)
	// ^按位异或（两位不一样则为1）
	fmt.Println(5 ^ 2)
	// <<将二进制左移指定位数
	fmt.Println(5 << 1)
	fmt.Println(1 << 10) // 1024
	// <<将二进制右移指定位数
	fmt.Println(5 >> 1)

	var m = int8(1)      // 只能存8位
	fmt.Println(m << 10) // 超出

	// 赋值运算符
	var x int
	x = 10
	x += 1  // x = x+1
	x -= 1  // x= x-1
	x *= 2  // x=x*2
	x /= 2  // x= x/2
	x %= 2  // x=x%2
	x <<= 2 // x=x<<2
	x &= 2  // x=x&2
	x |= 3  // x=x|2
	x ^= 4  // x= x^4
	x >>= 2 // x=x>>2
	fmt.Println(x)
}
