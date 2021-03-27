package main

import (
	"fmt"
)

func main()  {
	// 定义一个 map 变量
	cities := make(map[string]string)

	// 这里相当于添加，因为定义了 cities map 之后默认没有 key 的
	// 所以如果 key 没有，就是增加
	cities["no1"] = "北京"
	cities["no2"] = "重庆"
	cities["no3"] = "成都"

	fmt.Println(cities)

	// 这里我将 no3 修改为了 西安，因为 "no3" 这个 key 在上面是存在的
	// 所以当有存在 key 的时候就是对 value 的修改
	cities["no3"] = "西安"
	fmt.Println("没有删除之前的 map \n",cities)

	delete(cities,"no3")
	fmt.Printf("删除后的 map %v\n",cities)

	delete(cities,"no5")


	// map 查找
	// 查找 cities map 中是否存在 no2 这个 key
	val,ok := cities["no2"]
	if ok {
		fmt.Println("找到了 no2 这个key, val = ",val)
	} else {
		fmt.Println("没有找到 no2 这个 key")
	}

	// 一次性删除该 map 所有的 key
	// 1.遍历所有的 key 然后逐个删除
	// 2.直接使用 make 一个新的空间，利用 gc 回收机制进行删除。
	// 比如说我想将 cities 这个 map 清空就可以如下操作
	cities = make(map[string]string)
	fmt.Println("清空后的 cities",cities)

	

}