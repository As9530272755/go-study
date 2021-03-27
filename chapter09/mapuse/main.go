package main

import (
	"fmt"
)

func main()  {
	// 第一种使用方式
	var a map[string]string

	a = make(map[string]string,10)
	a["no1"] = "zgy"
	a["no2"] = "tom"
	a["no3"] = "jack"
	fmt.Println(a)
	fmt.Println()

	// 第二种方式
	var a1 = make(map[string]string)
	a1["no1"] = "a1-map"
	a1["no2"] = "a1-map1"
	a1["no3"] = "a1-map2"
	fmt.Println(a1)
	fmt.Println()

	// 第三种方式
	var a2 map[string]string = map[string]string {
		"no4" : "成都" ,
	}

	a2["no1"] = "北京"
	a2["no2"] = "重庆"
	fmt.Println(a2)

	// 第四种方式
	name := map[string]string {
		"name3" : "钢铁侠" ,
	}
	name["name1"] = "黄淑芬"
	name["name2"] = "张淑芳"
	fmt.Println(name)
}