package main

import (
	"fmt"
)

// 传统排序

// 数组是值类型，所以我在 BubbleSort 函数中通过 * 指向 arr 数组的一个指针
// 也就是说 BubbleSort 中的 arr 就是 main 中 arr
// 这样在 bubblesort 函数中进行的数组排序才会印象到 main 中的 arr 变量
func BubbleSort(arr *[5]int)  {
	fmt.Println("排序前 ",*arr)

	// 定义一个临时变量(用于交换)
	temp := 0

	// 完成第一轮排序结果：[24,69,57,13,80]
	// 第一轮只进行 4 次比较，所以循环次数就是 len(arr) -1 = 4
	for i := 0 ; i < len(arr) - 1 ; i++ {
		// 通过 if 进行判断，* 这里是指针，取 main 函数中的 arr 数组值，如果 arr[i] 次的值 大于 (*arr)[i +1]
		// 也就是说本次取得 arr 数组值大于下一次 arr 数组的值就在 if 体中进行交换
		if (*arr)[i] > (*arr)[i +1] {
			
			// 通过 temp 进行交换
			temp = (*arr)[i]
			(*arr)[i] = (*arr)[i +1]
			(*arr)[i +1] = temp
		}
	}
	// 输出第二次排序的结果
	fmt.Println("第一次排序 = ",*arr)

	// 第二轮排序
	for i := 0 ; i < len(arr) -2 ; i++ {
		if (*arr)[i] > (*arr)[i +1] {
			
			// 通过 temp 进行交换
			temp = (*arr)[i]
			(*arr)[i] = (*arr)[i +1]
			(*arr)[i +1] = temp
		}
	}
	// 输出第二次排序的结果
	fmt.Println("第二次排序 = ",*arr)

	// 第三次排序
	for i := 0 ; i < len(arr) -3 ; i++ {
		if (*arr)[i] > (*arr)[i +1] {
			
			// 通过 temp 进行交换
			temp = (*arr)[i]
			(*arr)[i] = (*arr)[i +1]
			(*arr)[i +1] = temp
		}
	}
	// 输出第三次排序的结果
	fmt.Println("第三次排序 = ",*arr)

	// 第4次排序
	for i := 0 ; i < len(arr) -4 ; i++ {
		if (*arr)[i] > (*arr)[i +1] {
			
			// 通过 temp 进行交换
			temp = (*arr)[i]
			(*arr)[i] = (*arr)[i +1]
			(*arr)[i +1] = temp
		}
	}
	// 输出第4次排序的结果
	fmt.Println("第四次排序 = ",*arr)
}


func main()  {
	// 定义数组
	var arr [5]int = [...]int{24,69,80,57,13}
	// 将数组传递个一个函数，完成排序
	// 调用 BubbleSort 函数传入 arr 数组的地址
	BubbleSort(&arr)
}