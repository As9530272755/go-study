package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

type CharCount struct {
	ChCount  int
	NumCount int
	Space    int
	Other    int
}

func (ch *CharCount) Count() {
	file, err := os.Open("d:/abc.txt")
	if err != nil {
		fmt.Print("open file err", err)
		return
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		str, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		for _, v := range str {
			switch {
			case v >= 'a' && v <= 'z':
				ch.ChCount++
			case v >= 'A' && v <= 'Z':
				ch.ChCount++
			case v >= '0' && v <= '9':
				ch.NumCount++
			case v == ' ' || v == '\t':
				ch.Space++
			default:
				ch.Other++
			}
		}
	}
	fmt.Printf("字母 = %v \n数字 = %v \n空格 = %v \n其他 = %v \n",
		ch.ChCount, ch.NumCount, ch.Space, ch.Other)
}

func main() {
	var charCount CharCount
	charCount.Count()
}
