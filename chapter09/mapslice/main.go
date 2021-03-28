package main

import (
	"fmt"
)

func main()  {
	// 演示 map 切片的使用
	// 1.定义一个 monster 的 map 切片,相当于说这个 monster 可以存放妖怪的信息，
	// 2.存放他的名字和年龄
	var monster []map[string]string

	// 3.在对切片进行make, 预计需要放入两个妖怪信息，monster 这个切片的类型为 map ，这里 2 也就是他的 len 的长度
	monster = make([]map[string]string,2)

	// 4.增加第一个妖怪的信息,因为 map 最开始的默认值为 nil(空)
	// 这里就是 monster[0] 如果为nil(空),就对 monster[0] 分配一个 map 信息。
	if monster[0] == nil {
		
		// 因为这个切片的数据类型是一个 map 所以我在使用的时候也需要 make 一下
		monster[0] = make(map[string]string)
		monster[0]["name01"] = "猪八戒"
		monster[0]["age01"] = "100岁"
	} 

	// 5.增加第二个妖怪信息，因为 map 最开始的默认值为 nil(空)
	// 这里就是 monster[1] 如果为nil(空),就对 monster[1] 分配一个 map 信息。
	if monster[1] == nil { 

		// 因为这个切片的数据类型是一个 map 所以我在使用的时候也需要 make 一下
		monster[1] = make(map[string]string)
		monster[1]["name02"] = "唐僧"
		monster[1]["age02"] = "200岁"
	}

	// 但是这里如果我还想动态添加一个妖怪信息，就会越界，所以这时候需要使用到切片中的 append 函数
	// 6.先定义一个新的 monster 信息,定义为 map 类型
	NewMonster := map[string]string{
		"name03" : "火云邪神",
		"age03" : "300岁",
	}

	// 这里我们再将 append 追加的内容传递给 monster 这个切片，不能传递给 NewMonster 因为 NewMonster 是一个 map 类型。
	monster = append(monster,NewMonster)

	fmt.Println(monster)
}