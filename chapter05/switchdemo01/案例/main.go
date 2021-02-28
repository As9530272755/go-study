package main

import (
	"fmt"
)

// func test(b) byte {
// 	return b + 1 // 此代码的意思是通过 test 函数穿了一个 b 进来然后在 b 的基础上 + 1
// }

func main() {
	// switch 'a' {
	// 	case 'a' :
	// 		fmt.Println("猴子穿新衣")
	// 	case 'b' :
	// 		fmt.Println("猴子没屁股")
	var age int = 10

	switch {
	case age == 10 :
		fmt.Println("age = 10")
	}
}