package main

import (
	"fmt"
)

func main()  {
	var s byte
	fmt.Println("请输入对应周几来吃对应的菜（如：1为周一）")
	fmt.Scanln(&s)
	switch s {
	case 1:
		fmt.Println("干煸豆角")
	case 2:
		fmt.Println("醋溜土豆")
	case 3:
		fmt.Println("红烧狮子头")
	case 4:
		fmt.Println("油炸花生米")
	case 5:
		fmt.Println("蒜蓉扇贝")
	case 6:
		fmt.Println("东北乱炖")
	case 7:
		fmt.Println("大盘鸡")
	default:
		fmt.Println("请输入正确的时间")
	}
}