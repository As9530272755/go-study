package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	file := "d:/xx.txt"
	file1, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", file1)
}
