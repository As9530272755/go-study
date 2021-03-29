package main

import (
	"fmt"
	"sort"
)

type Stu struct {
	Name string
	Age int
	Address string
}

func modify(map1 map[int]int)  {
	map1[5] = 50
}

func modifyUser(users map[string]map[string]string,names string)  {
	if users[names] != nil {
		users[names]["pwd"] = "888888"
	} else {
		users[names] = make(map[string]string)
		users[names]["names"] = names
		users[names]["pwd"] = "888888"
	}
}

func main()  {
	cities := make(map[string]map[string]string)

	cities["no1"] = make(map[string]string)
	cities["no1"]["alcohol"] = "泸州老窖"
	cities["no1"]["cities"] = "泸州"

	cities["no2"] = make(map[string]string)
	cities["no2"]["alcohol"] = "国宾"
	cities["no2"]["cities"]  = "重庆"

	for k,v := range cities {
		fmt.Printf("k = %v\n",k)
		for k1,v1 := range v {
			fmt.Printf("\tk = %v , v = %v\n",k1,v1)
		}
	}
	fmt.Printf("cities[no1] 中有 %v 对 key-value\n",len(cities["no1"]))

	var monster []map[string]string

	monster = make([]map[string]string,2)

	if monster[0] == nil {
		monster[0] = make(map[string]string)
		monster[0]["name01"] = "猪八戒"
		monster[0]["age"] = "100岁"
	}

	if monster[1] == nil {
		monster[1] = make(map[string]string)
		monster[1]["name"] = "唐山"
		monster[1]["age"] = "120岁"
	}
	
	NewMonster := make(map[string]string)
	NewMonster["name3"] = "火云邪神"
	NewMonster["age1"] = "200"
	NewMonster["name4"] = "孙悟空"
	NewMonster["age2"] = "800"

	monster = append(monster,NewMonster)

	fmt.Println(monster)

	for k,v := range cities {
		fmt.Printf("k = %v v = %v \n",k,v)
	}

	fmt.Println()

	map1 := make(map[int]int)
	map1[1] = 10
	map1[3] = 20
	map1[4] = 30
	map1[5] = 20

	for k,v := range map1 {
		fmt.Println(k,v)
	}

	fmt.Println()

	var keys[]int
	for k,_ := range map1 {
		keys = append(keys,k)
	}

	sort.Ints(keys)

	fmt.Println("排序后 keys = ",keys)

	for _,v := range keys{
		fmt.Printf("map1[%v] = %v\n",v,map1[v])
	}

	goods := make(map[string]map[string]string)
	goods["乐事薯片"] = make(map[string]string)
	goods["乐事薯片"]["出厂日期"] = "2021/03/29"
	goods["乐事薯片"]["重量"] = "200g"

	goods["卫龙辣条"] = make(map[string]string)
	goods["卫龙辣条"]["生产日期"] = "2020/12/30"
	goods["卫龙辣条"]["重量"] = "198g"
	
	goods["死神拉面"] = make(map[string]string)
	goods["死神拉面"]["生产日期"] = "2021/02/28"
	goods["死神拉面"]["重量"] = "300g"

	goods["自嗨锅"] = make(map[string]string)
	goods["自嗨锅"]["生产日期"] = "2021/01/29"
	goods["自嗨锅"]["重量"] = "400g"

	for k,v := range goods {
		fmt.Printf("商品名称 = %v \n",k)
		for k1,v1 := range v {
			fmt.Printf("商品属性 = %v 商品信息 = %v \n",k1,v1)
		}
	}

	fmt.Println()

	var goodsSlice[]string

	for k,_ := range goods {
		goodsSlice = append(goodsSlice,k)
	}

	sort.Strings(goodsSlice)

	fmt.Println(goodsSlice)

	for i,v := range goodsSlice {
		fmt.Printf("商品排序 = %v 信息 = %v 具体信息 = %v\n",i,v,goods[v])
	}

	modify(map1)
	fmt.Println(map1)

	students := make(map[string]Stu)
	stu1 := Stu{"张三",20,"北京"}
	stu2 := Stu{"黄",30,"重"}

	students["no1"] = stu1
	students["no2"] = stu2
	fmt.Println(students)
	fmt.Println()

	for k,v := range students {
		fmt.Println("编号",k)
		fmt.Println("姓名",v.Name)
		fmt.Println("年龄",v.Age)
		fmt.Println("地址",v.Address)
		fmt.Println()
	}

	user := make(map[string]map[string]string)
	user["张桂元"] = make(map[string]string)
	user["张桂元"]["pwd"] = "123"

	modifyUser(user,"张桂元")
	modifyUser(user,"tom")
	fmt.Println(user)
}