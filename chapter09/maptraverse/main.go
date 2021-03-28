package main

import (
	"fmt"
)

func main()  {
	// 使用 for-range 遍历 map
	
	// 一个相对简单的 map 来进行遍历
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "重庆"

	fmt.Printf("cities 有 %v 对 key-value\n",len(cities))
	// 使用 for-range 的方式将每一对 key-value 遍历出来
	// 这里的 k 是他的 key ，v 是每一个 k 对应的 value
	for k,v := range cities {
		fmt.Printf("k = %v , v = %v\n",k,v)
	}

	// 对相对复杂的 map 进行遍历
	// 我们可以看到这个 map 他的值也是一个 map 类型
	// 所以我们在遍历的时候需要对他进行两次遍历，也就是双层 for-range
	studentMap := make(map[string]map[string]string)
	studentMap["stu01"] = make(map[string]string)
	studentMap["stu01"]["name"] = "tom"
	studentMap["stu01"]["sex"] = "男"
	studentMap["stu01"]["address"] = "重庆"

	studentMap["stu02"] = make(map[string]string)
	studentMap["stu02"]["name"] = "zgy"
	studentMap["stu02"]["sex"] = "男"
	
	studentMap["stu03"] = make(map[string]string)
	studentMap["stu03"]["name"] = "黄二狗"
	studentMap["stu03"]["sex"] = "女"

	fmt.Printf("studentMap 有 %v 对 key-value\n",len(studentMap))
	// 使用 for-range 遍历一个比较复杂的 map
	// 第一次 for-range 中的 k1 就是 stu01、stu02、stu03 这样的 key ，然后这个 v1 他又是一个 map 
	for k1,v1 := range studentMap{
		fmt.Println("k1 = ",k1)
		// 第二次就要对 v1 这个 map 进行一个 for-range，这里面的 k2 就是对应的 "name"、"sex"、"address" 这样的 key
		// v2 就是对应的 "tom".... 这样具体的值
		for k2,v2 := range v1 {
			fmt.Printf("\t k2 = %v v2 = %v\n",k2,v2)
		}
	}

}