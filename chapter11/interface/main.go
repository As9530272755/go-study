package main

import (
	"fmt"
)

// 声明一个接口
type Usb interface {
	
	// 声明了两个没有实现的方法
	Start()
	Stop()
}

type Usb2 interface {
	Start()
	Stop()
}

type Phone struct {

}

// 让 Phone 结构体实现 Usb 接口的两个方法
func (p Phone) Start() {
	fmt.Println("手机 Usb2 开始工作...")
}

func (p Phone) Stop() {
	fmt.Println("手机 Usb2 停止工作...")
}

// 让相机结构体也实现 Usb 接口的方法
type Camera struct {

}

func (c Camera) Start() {
	fmt.Println("相机开始工作...")
}

func (c Camera) Stop() {
	fmt.Println("相机停止工作...")
}

// 让电脑结构体实现 Usb 接口的方法
type Computer struct {

}

// 编写一个 working 方法,usb 类型是 Usb 接口类型
// 也就是 working 这里可以接收一个 Usb 接口类型变量
// 只要实现了 Usb 接口，(所谓实现 usb 接口，就是指实现了 Usb 接口声明的所有方法)
func (c Computer) Working(usb Usb2) {
	
	// 这里通过 usb 接口来调用 Start 和 Stop 方法
	usb.Start()
	usb.Stop()
}

func main()  {
	// 测试
	// 声明结构体变量
	computer := Computer{}
	phone := Phone{}
	camera := Camera{}

	// 关键点,调用 computer 结构体中的 Working 方法，分别传入 phone 和 camera 结构体
	computer.Working(phone)
	computer.Working(camera)
}