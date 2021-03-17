package main

import (
	"fmt"
)

func main()  {
	// 用 append 内置函数，可以对切片进行动态增加

	// 定义一个切片,并且给他了 100，200，300 这么三个数组
	var slice []int = []int{100,200,300}

	// 通过第一次输出 slice 切片得到的数组就是 100，200，300
	fmt.Println("第一次输出 slice = ",slice)

	// 但是这次我想追加一个 400 到 slice 切片中，就可以通过 append 追加具体的元素
	slice = append(slice,400)
	fmt.Println("第一次追加后 slice = ",slice)

 
	// 假如一次追加不够我们还可以接着使用 append 继续追加，而且可以一次性追加多个元素
	slice = append(slice,500,600)
	fmt.Println("多次追加后 slice = \n",slice)

	// 通过 append 追加切片给 slice
	slice = append(slice,slice...)
	fmt.Println("把切片在追加 slice = \n",slice)
}