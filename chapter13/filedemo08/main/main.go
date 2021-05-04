package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	filePath := "f:/kkk.txt"
	file, err := os.Open(filePath)

	if err != nil {
		fmt.Printf("open file err = %v and create %v\n", err, filePath)
		file1, err := os.OpenFile(filePath, os.O_CREATE, 0666)
		if err != nil {
			fmt.Printf("create file err = %v\n", err)
		}
		defer file1.Close()
	}

	defer file.Close()

	file1Path := "d:/xx.txt"
	data, err := ioutil.ReadFile(file1Path)
	if err != nil {
		fmt.Printf("Read file err = %v\n", err)
		return
	}

	err = ioutil.WriteFile(filePath, data, 0666)
	if err != nil {
		fmt.Printf("Write file err = %v\n", err)
	}
}
