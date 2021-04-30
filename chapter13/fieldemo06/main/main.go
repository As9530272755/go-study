package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "d:/1122.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Print("open file err", err)
		return
	}

	defer file.Close()

	str := "好好学习!天天向上!\n"
	writer := bufio.NewWriter(file)
	writer.WriteString(str)
	writer.Flush()
}
