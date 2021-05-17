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
	fmt.Println(p.Name, "手机启动中...")
}

func (p Phone) Stop() {
	fmt.Println(p.Name, "手机停止中...")
}

func (p Phone) Call() {
	fmt.Println(p.Name, "拨打电话中...")
}

type Camera struct {
	Name string
}

func (c Camera) Start() {
	fmt.Println(c.Name, "相机启动中...")
}

func (c Camera) Stop() {
	fmt.Println(c.Name, "相机停止中...")
}

type Tool struct {
}

func (tool Tool) working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main() {
	phone := Phone{}

	var UsbArr [2]Usb
	tool := Tool{}

	UsbArr[0] = Phone{
		Name: "小米",
	}
	UsbArr[1] = Camera{
		Name: "佳能",
	}

	for i := 0; i < len(UsbArr); i++ {
		tool.working(UsbArr[i])
		phone.Call()
	}
}
