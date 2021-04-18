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

func (p *Phone) Start() {
	fmt.Println("手机正在启动...")
}

func (p *Phone) Stop() {
	fmt.Println("手机已经停止...")
}

type Camera struct {

}

func (c *Camera) Start() {
	fmt.Println("相机正在启动...")
}

func (c *Camera) Stop() {
	fmt.Println("相机已经停止...")
}

type Tool struct {

}

func (t *Tool) Working(usb Usb) {
	usb.Start()
	usb.Stop()
}

func main()  {
	phone := Phone{}
	tool := Tool{}
	camera := Camera{}
	tool.Working(&phone)
	tool.Working(&camera)
}