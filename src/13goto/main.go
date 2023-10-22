package main

import "fmt"

// goto
func main() {
	var flag bool
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				flag = true
				break
			}
			fmt.Printf("%v-%c\n", i, j)
		}
		if flag {
			break
		}
	}

	fmt.Println("----------")
	// goto + label实现跳出多层for循环
	for i := 0; i < 10; i++ {
		for j := 'A'; j < 'Z'; j++ {
			if j == 'C' {
				goto XX
			}
			fmt.Printf("%v-%c\n", i, j)

		}
	}
XX:
	fmt.Println("over")
}
