package main

import (
	"fmt"
	"os"
	"bufio"
	"io"
)

func main()  {
	file,err := os.Open("d:/xx.txt")
	if err != nil {
		fmt.Println("open file err = ",err)
	}

	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		
		file,err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Print(file)
	}
}