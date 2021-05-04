package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {
	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Print("open file err", err)
	}
	defer srcFile.Close()

	readFile := bufio.NewReader(srcFile)

	dstFile, err := os.OpenFile(dstFileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print("Create file err", err)
	}
	defer dstFile.Close()

	wrtierFile := bufio.NewWriter(dstFile)

	return io.Copy(wrtierFile, readFile)
}

func main() {
	srcFileName := "d:/7.jpg"
	dstFileName := "f:/7.png"
	_, err := CopyFile(srcFileName, dstFileName)
	if err == nil {
		fmt.Print("copy file success")
	} else {
		fmt.Print("copy file fail")
	}
}
