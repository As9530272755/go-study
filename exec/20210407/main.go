package main

import (
	"fmt"
)

type Person struct {
	Name string
	Age int
	Scores [5]float64
	Ptr *int
	Slice []int
	Map1 map[string]map[string]string
}

func main()  {
	var p1 Person
	p1.Slice = make([]int,10,10)
	p1.Slice[0] = 100
	fmt.Println(p1.Slice,len(p1.Slice),cap(p1.Slice))

	p1.Map1 = make(map[string]map[string]string)
	p1.Map1["no1"] = make(map[string]string)
	p1.Map1["no1"]["name"] = "zgy"
	p1.Map1["no1"]["age"] = "10岁"
	fmt.Println(p1.Map1)
}