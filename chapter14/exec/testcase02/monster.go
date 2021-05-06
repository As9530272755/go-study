package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Monster struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func (monster *Monster) Store() bool {
	monster.Name = "牛魔王"
	monster.Age = 100
	monster.Skill = "牛魔拳"

	// json 序列化 monster 结构体
	data, err := json.Marshal(&monster)
	if err != nil {
		fmt.Println("open json err = ", err)
		return false
	}

	// 需要创建的文件路径
	filePath := "d:/monster.txt"
	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Print("create file err = ", err)
		return false
	}

	defer file.Close()

	// 将 data 写入到 file 文件中
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))
	writer.Flush()
	return true
}

func (monster *Monster) ReStore() bool {

	// 读取文件
	filePath := "d:/monster.txt"
	data, err := ioutil.ReadFile(filePath)
	if err == io.EOF {
		fmt.Print("reader file err = ", err)
		return false
	}
	//fmt.Printf("%s", data)

	err = json.Unmarshal(data, monster)
	if err != nil {
		fmt.Print("json open err", err)
		return false
	}
	return true
}
