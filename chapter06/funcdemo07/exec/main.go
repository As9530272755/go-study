package main

import (
	"fmt"
	"strings"
)

// 1)自定义编写一个函数 makeSuffix(suffix string) 可以接收一个文件后缀名(比如.jpg)，并返回一个闭包。
// 2)调用闭包，可以传入一个文件名，如果该文件名没有指定的后缀(比如.jpg) ,则返回文件名.jpg ,如果已经有.jpg后缀，则返回原文件名。
// 3)要求使用闭包的方式完成。
// 4) strings.HasSuffix 该函数可以判断某个字符串是否有指定的后缀 

func makeSuffix(suffix string) func (string) string {
	return  func (name string) string {
		// 如果 name 没有指定后缀，则加上.jpg 后缀，否则就返回原来的名字
		if !strings.HasSuffix(name,suffix) {
			return name + suffix
		} 
		return name
	}
}

func main()  {
	f := makeSuffix(".jpg")
	// t := makeSuffix
	// f1 := t(".txt")
	fmt.Println("文件名处理后为",f("winter")) //f("winter") = f := makeSuffix(".jpg") winter winter 是一个返回值
	// fmt.Println("f1=",f1("zgy"))
} 