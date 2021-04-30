package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	filePath := "d:/xx.txt"

	// 使用 os.O_RDWR 和 os.O_APPEND 对文件执行读写和追加操作
	file, err := os.OpenFile(filePath, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open file err", err)
	}

	// defer 函数最后关闭文件句柄
	defer file.Close()

	// 对文件
	read := bufio.NewReader(file)
	for {
		str, err := read.ReadString('\n')
		if err == io.EOF { // 如果读取到文件末尾就不处理
			break
		}
		// 显示到终端，不用换行显示
		fmt.Print(str)
	}

	// 写文件
	str := "hello 重庆\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString(str)
	}
	writer.Flush()
}
