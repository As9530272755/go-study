package main

import (
	"fmt"
)

// 定义 Usb 接口，方法为 Say
type Usb interface {
	Say()
}

// 定义 Stu 结构体
type Stu struct {

}

// 给 Stu 结构体绑定方法 Say
func (this *Stu) Say() {
	fmt.Println("Say()...")
}

func main()  {
	// 定义 stu 变量类型为 Stu 结构体
	var stu Stu = Stu{}

	// 定义 u 变量类型为 Usb 接口，传递 stu 结构体的地址
	var u Usb = &stu
	
	// 调用 u 接口中的 Say 方法
	u.Say()
}