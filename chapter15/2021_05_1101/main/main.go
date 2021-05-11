package main

import (
	"bufio"
	"fmt"
	"os"
)

type Hobby struct {
	Game    string
	Smoke   string
	Alcohol string
}

func (hobby *Hobby) ShowHobby() string {
	str := fmt.Sprintf("给李岚洁的以下几点要求：\n1.要阳光：%v\n2.要快乐：%v\n3.要积极向上：%v\n",
		hobby.Game, hobby.Smoke, hobby.Alcohol)
	return str
}

type Address struct {
	WorkPlace     string
	AncestralHome string
}

func (address *Address) ShowAddress() string {
	str := fmt.Sprintf("希望!\n%v\n%v\n",
		address.WorkPlace, address.AncestralHome)
	return str
}

type Info struct {
	Age  int
	Name string
	Sex  string
}

func (info Info) ShowInfo() string {
	str := fmt.Sprintf("年龄：%v\n姓名：%v\n性别：%v\n",
		info.Age, info.Name, info.Sex)
	return str
}

type Person struct {
	Hobby
	Address
	Info
}

func OpenFile() {
	person := Person{
		Hobby{Game: "愿你如冬日里的那一丝丝暖阳，不仅能温暖别人还要温暖自己！！", Smoke: "快乐似风，抓不住追不到，但是你能感受到！", Alcohol: "不为别人，只为自己加油，祝你走出阴霾"},
		Address{WorkPlace: "我们总会说一些令人后悔的事\n每个字都是那么刺耳\n我们只需要冷静下来\n现在一个人出去静静\n事情总会好起来的\n不要气馁", AncestralHome: "Come on, I believe you"},
		Info{Age: 25, Name: "曦泽", Sex: "男"},
	}
	str := "笑靥如花 风姿绰约\n"

	filePath := "c:/李岚洁.txt"

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file err = ", err)
	}

	defer file.Close()

	write := bufio.NewWriter(file)
	write.WriteString(str)
	write.WriteString(person.ShowAddress())
	write.WriteString(person.ShowHobby())
	// write.WriteString(person.ShowInfo())

	err = write.Flush()
	if err != nil {
		fmt.Println("write file err ", err)
	}

}

func main() {
	OpenFile()
}
