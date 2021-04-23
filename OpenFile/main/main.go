package main

import (
	"os"
	"io"
	"fmt"
)

// 该函数读取文件 128 个字节类容
func readFromFile()  {
	// fileObj 是打开文件对象，err 是返回他的一个错误
	fileObj,err := os.Open("xx.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	
	defer fileObj.Close()	// 程序即将退出的时候关闭文件

	// 读取文件内容,make 一个切片，长度为 128 个字节
	tmp := make([]byte,128)
	n, err := fileObj.Read(tmp)
	if err != nil {
		fmt.Printf("read form file,err:%v\n",err)
		return
	}

	fmt.Printf("read %d bytes from file,\n",n)
	fmt.Println(string(tmp))
}

// 读取文件所有内容
func readAll()  {
	fmt.Println("readAll~~~")
	fmt.Println("readAll~~~~~")

	// fileObj 是打开文件对象，err 是返回他的一个错误
	fileObj,err := os.Open("xx.txt")
	if err != nil {
		fmt.Printf("open file failed,err:%v\n",err)
		return
	}
	
	defer fileObj.Close()	// 程序即将退出的时候关闭文件


	for {
		// 读取文件内容,make 一个切片，长度为 128 个字节
		tmp := make([]byte,128)
		n, err := fileObj.Read(tmp)
		
		// 通过 if err 来判断该文件是否不能再读，即退出该程序
		if err == io.EOF{
			// 把当前读了多少个字节的数据打印出来，然后退出
			fmt.Println(string(tmp[:n]))
			return
		}
		if err != nil {
			fmt.Printf("read form file,err:%v\n",err)
			return
		}

		fmt.Printf("read %d bytes from file,\n",n)
		fmt.Println(string(tmp[:n]))
	}
}

// 文件操作
func main()  {
	readFromFile()
	readAll()
}