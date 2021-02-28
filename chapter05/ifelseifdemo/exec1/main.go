package main

import (
	"fmt"
)

func main()  {
	var height   int32
	var money    float32
	var handsome bool
	fmt.Println("请输入身高（cm）：")
	fmt.Scanln(&height)
	fmt.Println("请输入财富：")
	fmt.Scanln(&money)
	fmt.Println("如果帅请输入 true 否则输出 false：")
	fmt.Scanln(&handsome)
	if height >= 180 && money >= 10000000 && handsome == true {
		fmt.Println("我一定要嫁给他!!!")
	} else if height >= 180 || money >= 10000000 || handsome == true {
		fmt.Println("嫁吧，比上不足，比下有余。")
	} else {
		fmt.Println("滚不嫁！！")
	}
}