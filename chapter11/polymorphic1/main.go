package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	Name string
}

func (p Phone) Start() {
	fmt.Println("手机正在启动...")
}

func (p Phone) Stop() {
	fmt.Println("手机已经停止...")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println("相机正在启动...")
}

func (c Camera) Stop() {
	fmt.Println("相机已经停止...")
}

func main()  {
	// 定义一个 usb 接口数组，可以存放 phone 和 camera 结构体变量
	// 这里就体现出多态数组

	// 定义 usbArr 数组有 3 个元素，类型为 Usb 接口
	var usbArr[3]Usb

	// 此时在 usbArr 数组中已经存放了两种结构体了，一种是 phone 一种是 camera
	// 如果 usbArr 数组不是接口类型是绝对放不进去的，这里是利用他多态的特点
	usbArr[0] = Phone{"vivo-手机"}
	usbArr[1] = Phone{"小米-手机"}
	usbArr[2] = Camera{"索尼-相机"}
	fmt.Println(usbArr)
}