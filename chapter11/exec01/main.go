package main

import (
	"fmt"
	"sort"
	"math/rand"
)

// 1.定义 Hero 结构体
type Hero struct {
	Name string
	Age int
}

// 2.声明 Hero 结构体切片类型
type HeroSlice []Hero

// 3.实现 Interface 接口
func (hs HeroSlice) Len() int {
	return len(hs)
}

// Less 这个方法就是决定使用什么标准进行排序
// 1.按 Hero 的Name从小到大排序
func (hs HeroSlice) Less(i,j int) bool {
	return hs[i].Name < hs[j].Name
}

// 定义 Swap 方法
func (hs HeroSlice) Swap(i,j int) {
	temp := hs[i]
	hs[i] = hs[j]
	hs[j] = temp
}

func main()  {
	// 先定义一个切片
	var intSlice = []int{12,-2,10,7,9,90}
	
	// 对 intSlice 进行排序
	// 使用冒泡算法对切片进行排序
	temp := 0
	for i := 0 ; i < len(intSlice) -1; i++ {
		for j := 0 ; j < len(intSlice) -1; j++ {
			if intSlice[j] > intSlice[j+1] {
				temp = intSlice[j+1]
				intSlice[j+1] = intSlice[j]
				intSlice[j] = temp
			}
		}

	}
	fmt.Println("冒泡算法对切片进行排序：",intSlice)

	// 使用系统提供的方法对切片进行排序
	var Slice = []int{0,-1,10,7,90,20}
	sort.Ints(Slice)
	fmt.Println("使用 sort 包对切片进行排序",Slice)

	// 对一个结构体切片进行排序
	// 1.系统提供 Sort 包的方法...
	// 测试对结构体接片进行排序
	var heroes HeroSlice
	for i := 0 ; i < 10 ;i++ {
		hero := Hero {
			Name : fmt.Sprintf("英雄~%d",rand.Intn(100)),
			Age : rand.Intn(100),
		}
		// 将 hero append 到 heroes 切片中
		heroes = append(heroes,hero)
	}
	// 排序前的顺序
	for _,v := range heroes {
		fmt.Println("对 hero 切片结构体排序前：",v)
	}

	// 调用 Sort 方法
	sort.Sort(heroes)
	fmt.Println("排序后：")
	for _,v := range heroes {
		fmt.Println("对 hero 结构体排序后：",v)
	}
}