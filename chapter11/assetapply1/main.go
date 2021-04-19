package main

import (
	"fmt"
)

type Usb interface {
	Start()
	Stop()
}

type Phone struct {

}

func (p Phone) Start() {
	fmt.Println("手机正在启动...")
}

func (p Phone) Stop() {
	fmt.Println("手机正在停止...")
}

func (p Phone) Call() {
	fmt.Println("正在拨打电话..")
}

type Camera struct {

}

func (c Camera) Start() {
	fmt.Println("相机正在启动...")
}

func (c Camera) Stop() {
	fmt.Println("相机正在停止...")
}

type Tool struct {

}

func (t Tool) Working(usb Usb) {
	usb.Start()
	usb.Stop()
	phone,ok := usb.(Phone)
	if ok {
		phone.Call()
	}
}

func main()  {
	var usbArr[2]Usb
	usbArr[0] = Phone{}
	usbArr[1] = Camera{}
	tool := Tool{}
	for _,v := range usbArr{
		tool.Working(v)
		fmt.Println()
	}
}