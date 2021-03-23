package main

import (
	"fmt"
)

/*
思路：
1 如果 arr[middle] > findVal ： 就应该向 leftindex 到 middle -1 这个区间进行查找 ；
2 如果 arr[middle] < findVal： 就应该向 middle + 1 到 rightindex 这个区间进行查找；
3 如果 arr[middle] = findVal：就找到了；
*/

// 定义二分查找的函数，调用 main 函数中的 arr 数组，
func BinaryFind(arr *[6]int,leftIndex int,rightIndex int,findVal int)  {
	
	// 判断 leftIndex 是否大于 rightIndex
	if leftIndex > rightIndex{
		fmt.Println("找不到")
		return
	}
	
	// 先找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	
	// 通过 if 判断 middle 是否大于 findVal 如果为真就表示我们想要找的值就应该在  leftindex 到 middle - 1 之间
	// 并且通过 * 取 arr 数组的指针值
	if (*arr)[middle] > findVal {
		// 这个时候传入 arr 就不用指针符了，因为这个 arr 就是 BinaryFind 函数的 arr
		BinaryFind(arr,leftIndex,middle - 1,findVal)
		
	// else if 判断 middle 是否小于 findVal 如果为真就表示我们想要找的值就应该在  middle + 1 到 rightIndex 之间
	} else if (*arr)[middle] < findVal {
		BinaryFind(arr,middle + 1,rightIndex,findVal)
	} else {
		// 还有一种可能性就是 arr[middle] = findVal 就说明找到了
		fmt.Printf("找到了，下标为： %v\n",middle)
	}
}

func main()  {
	// 定义一个从小到大的数组
	var arr [6]int = [...]int{1,8,10,89,1000,1234}

	// 测试 
	// 将 main 函数中 arr 通过 & 的方式传入地址，
	// 0 表示传递给了 BinarFind 函数中的 leftIndex 形参，也就是 leftIndex 下标
	// len(arr) -1 表示传递给了 BinarFind 函数中的 rightindex 形参，也就是 rightindex 下标
	// 这里我查找 1000，也就是传递给了 BinarFind 函数中的 findVal 形参
	BinaryFind(&arr,0,len(arr)-1,1000)
}