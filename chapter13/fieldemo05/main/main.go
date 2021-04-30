package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "d:/xx.txt"

	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		fmt.Print("open file err", err)
	}

	defer file.Close()

	str := "今天是个好日子!!\n"

	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString(str)
	}

	writer.Flush()

}
