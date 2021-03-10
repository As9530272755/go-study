package main

import (
	"fmt"
)

// [3]int 在 go 中会是一种数据类型，而 [4]int 又是另外一种数据类型，因为他们数组之间的长度不一样
func test01(arr [3]int)  { 
	// 在 test01 函数里面讲 arr 数组的第一个元素改成了 88
	arr[0] = 88
	fmt.Println("test01() = ",arr)
}

// 这里我加了一个 * 在 [3]int 这里，这时候 arr 就变成了指向数组的一个指针
func test02(arr *[3]int)  {
	
	// 这里再将 arr 用 () 包起来，并且里面用 * 作为一个取值符，这样就能够取到这个数组了 
	// 因为 arr 这里本身就是一个指针，用了取值符就能够取到这个数组了
	(*arr)[0] = 99

}

func main()  {
	
	var arr01 [3]int;
	arr01[0] = 1
	arr01[1] = 30
	arr01[2] = 1
	fmt.Println(arr01)

	// 数组创建以后，如果没有赋值，是有默认值（默认值为 0）
	// 数值类型数组：默认值为 0
	// 字符串数组：默认值为 ""（就是空）
	// bool数组：默认值为 false

	var arr02 [3]float32
	var arr03 [3]string
	var arr04 [3]bool

	fmt.Printf("arr02 数值默认值 = %v\n",arr02)
	fmt.Printf("arr03 字符串默认值 = %v\n",arr03)
	fmt.Printf("arr04  布尔值默认值 = %v\n",arr04)

	//Go 的数组是属于值类型，在默认情况下是值传递，因此会进行值拷贝。数组间不会相互影响
	arr := [3]int{11,22,33}
	test01(arr)
	fmt.Println("main () = ",arr)

	arr1 := [3]int{44,55,66}
	test02(&arr1)
	fmt.Println("main() =",arr1)
}