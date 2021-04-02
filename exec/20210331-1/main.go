package main

import (
	"fmt"
	"sort"
)

func main()  {
	ints := make(map[string]map[string]string)
	ints["n1"] = make(map[string]string)
	ints["n1"]["name1"] = "大"
	ints["n1"]["sex1"] = "M"

	ints["n3"] = make(map[string]string)
	ints["n3"]["name1"] = "西"
	ints["n3"]["sex1"] = "M"

	ints["n2"] = make(map[string]string)
	ints["n2"]["name1"] = "话"
	ints["n2"]["sex1"] = "M"

	ints["n4"] = make(map[string]string)
	ints["n4"]["name1"] = "游"
	ints["n4"]["sex1"] = "M"

	var sort1 []string

	for k,_ := range ints {
		sort1 = append(sort1,k)
	}

	sort.Strings(sort1)
	for _,v := range sort1 {
		fmt.Printf("key 编号为 %v 值 = %v \n",v,ints[v])
	}

}