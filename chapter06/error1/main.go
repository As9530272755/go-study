package main

import (
	"fmt"
	"errors"
)

// 定义一个函数，该函数去读取配置文件 config.ini 的信息，
// 如果文件名传入不正确我们就抛出一个之定义的错误
func readConf(name string) (err error) {
	// 这里判断如果 name 等于 config.ini 就返回 nil 
	if name == "config.ini" {
		// 读取...
		return nil

	// 否则就调用 errors 包中的 New 函数并且输出 读取问文件错误
	} else {
		// 返回一个自定义的错误，将这个错误传给 test02 函数中的 err 
		return errors.New("读取文件错误。。")
	}
}

func test02()  {
	// 调用 readconf 函数，并且传入 name 变量，这里我们通过手动输入文件名
	var name string
	fmt.Println("请输入文件名")
	fmt.Scanln(&name)
	err := readConf(name) 
	
	// 如果读取文件发生错误，就输出这个错误，并且终止该程序
	if err != nil {
	
	// 如果当 err 不等于 nil 是就会执行 panic 然后将这个错误输出，并且程序终止也不会输出下面得代码
		panic(err)
	}

	// 如果文件名等于 config.ini 就继续执行下面的代码
	fmt.Println("test02() 继续执行")
}

func main()  {
	// 测试自定义错误的使用，调用 test02 函数
	test02()
	fmt.Println("test02() 继续执行")
}