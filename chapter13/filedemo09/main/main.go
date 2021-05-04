package main

import (
	"fmt"
	"os"
)

func PathExists(path string) (bool, error) {
	// 调用 os.Stat 函数将 path 变量传入，判断该 path 是否存在，如果存在返回 true
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}

	// 假如 path 不存在则通过 os.IsNotExist 函数判断，false 则索命这个文件或者这个目录不存在
	if os.IsNotExist(err) {
		return false, nil
	}

	// 如果为假同时还有个 err 就为其他信息
	return false, err
}

func CreateFile(filePath string, bool1 bool) (file_err string) {
	file_err = ""

	if !bool1 {
		file, err := os.OpenFile(filePath, os.O_CREATE, 0666)
		if err != nil {
			file_err = "Create file err"
		} else {
			file_err = "Create file success"
		}
		defer file.Close()
	} else {
		fmt.Println("file is exist")
	}
	return file_err
}

func main() {
	filePath := "f:/zgy.txt"
	bool1, _ := PathExists(filePath)
	createInfo := CreateFile(filePath, bool1)
	fmt.Println(createInfo)
}
