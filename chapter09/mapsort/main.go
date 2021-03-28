package main

import (
	"fmt"
	"sort"
)

func main()  {
	// map 的排序
	map1 := make(map[int]int,10)
	map1[10] = 100
	map1[1] = 13
	map1[2] = 4
	map1[4] = 90
	map1[6] = 10

	for k,v := range map1 {
		fmt.Printf("排序前 key = %v value = %v\n",k,v)
	}

	// 换行
	fmt.Println()

	// 如何按照 map 的 key 的顺序进行排列输出
	// 1.先将 map 的 key 放入到一个切片中
	// 2.对这个切片进行排序
	// 3.遍历这个切片，然后按照这个 key 来输出 map 的值

	// 定义一个切片,因为上面 map 类型为 int ，所以这个切片类型也为 int
	var keys []int

	// 对 map1 进行 for-range 遍历
	for k,_ := range map1{

		// 然后将这个 k 也就是锁 map1 的 key 通过 append 追加到 keys 这个切片中
		keys = append(keys,k)
	}

	// 对 keys 这个切片进行排序,这里使用的 sort 包中的 Ints 函数
	sort.Ints(keys)

	// 然后将 keys 这个切片进行输出
	fmt.Println("排序后 keys = ",keys)

	// 换行
	fmt.Println()

	// 遍历切片，将其输出对应的 key 值和 map 值
	for _,v := range keys{
		fmt.Printf("map[%v] = %v \n",v,map1[v])
	}
}