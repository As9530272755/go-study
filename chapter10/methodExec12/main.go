package main

import (
	"fmt"
)

type Box struct {
	Len int
	Wide int
	High int
}

func (box Box) info() int {
	volume := box.Len * box.Wide * box.High
	return volume
}

func main()  {
	var b Box
	fmt.Println("请输入长度:")
	fmt.Scanln(&b.Len)
	
	fmt.Println("请输入宽度:")
	fmt.Scanln(&b.Wide)

	fmt.Println("请输入高度:")
	fmt.Scanln(&b.High)

	volume := b.info()
	fmt.Println("该立方体体积为",volume)
}