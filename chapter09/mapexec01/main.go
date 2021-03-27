package main

import (
	"fmt"
)

func main()  {
	// 我们要存放 3 个学生信息，每个学生有 name 和 sex（性别） 信息
	studentMap := make(map[string]map[string]string)
	studentMap["no1"] = make(map[string]string,3) // 在使用多重 map 的时候 make(map[string]string 这句话绝对不能少
	studentMap["no1"]["name"] = "Tom"
	studentMap["no1"]["sex"] = "男"
	studentMap["no1"]["address"] = "北京长安街"

	studentMap["no2"] = make(map[string]string) // 在使用多重 map 的时候 make(map[string]string 这句话绝对不能少
	studentMap["no2"]["name"] = "zgy"
	studentMap["no2"]["sex"] = "男"

	studentMap["no3"] = make(map[string]string) // 在使用多重 map 的时候 make(map[string]string 这句话绝对不能少
	studentMap["no3"]["name"] = "Jack"
	studentMap["no3"]["sex"] = "女"

	fmt.Println(studentMap["no1"])
	fmt.Println(studentMap["no2"])
	fmt.Println(studentMap["no3"])

	fmt.Printf("第二个学生的姓名 %v\n",studentMap["no2"]["name"])

	fmt.Printf("第二个学生的性别 %v",studentMap["no2"]["sex"])



}