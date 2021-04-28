package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "d:/1122.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print("open file err", err)
	}

	defer file.Close()

	str := "卢本伟NB\n"
	writer := bufio.NewWriter(file)

	for i := 0; i < 3; i++ {
		writer.WriteString(str)
	}

	writer.Flush()
}
