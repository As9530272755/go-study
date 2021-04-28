package main

import (
	"fmt"

	// 因为要对文件进行操作所以需要 os 包
	"os"
)

func main() {

	// 打开一个文件
	// 概念说明：关于 file 的叫法（下面三种概念都是一个意思）
	// 1.file 叫 file 对象
	// 2.file 叫 file 指针
	// 3.file 叫 file 文件句柄
	file, err := os.Open("d:/xx.txt")
	if err != nil {
		fmt.Println("open file error = ", err)
	}

	// 如果打开成功了输出文件内容看看文件
	fmt.Println("file = ", file)

	// 当我们使用完了之后需要关闭文件
	// 需要使用 close 方法， close 方法是和 file 结构体进行绑定的一个方法
	// 并且该方法也会返回 error 错误，如果这个 Close 返回错误就会将返回值赋值给 err
	err = file.Close()
	if err != nil {
		fmt.Println("close file error = ", err)
	}
}
