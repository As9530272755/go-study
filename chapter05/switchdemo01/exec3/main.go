package main

import (
	"fmt"
)

func main()  {
	var season byte
	fmt.Println("输出对应的季节（春季：3、4、5）（夏季：6、7、8）（秋季：9、10、11）（冬季：12、1、2）")
	fmt.Scanln(&season)
	switch season {
		case 3,4,5 :
			fmt.Println("春")
		case 6,7,8 :
			fmt.Println("夏")
		case 9,10,11 :
			fmt.Println("秋")
		case 12,1,2 :
			fmt.Println("冬")
		default:
			fmt.Println("输入正确的季节")
	}
}