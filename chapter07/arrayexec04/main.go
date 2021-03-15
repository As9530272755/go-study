package main

import (
	"fmt"
	"math/rand"	// 引入 math/rand 包
	"time"		// 引入 time 包
)

func main()  {
	// 要求：随机生成 5 个数，并将其反转打印，复杂应用
	// 思路：
	// 1.随机生成 5 个数，在 golang 中需要使用到 rand.Intn 的这个函数
	// 2.当我们得到随机数后，就放到一个数组中 int 类型数组即可
	// 3.要求反转打印，将后面的打印到前面去，前面的打印到后面，交换的次数是 数组的大小除以 2
	// 4.交换的规律是第一个元素和最后一个元素进项交换，第二个元素和倒数第二个元素进行交换，以此类推

	var intArr [5]int
	// 为了每次生成的这个随机数不一样，我们需要给一个 seed 值，这里是 unixNano 纳秒秒数时间戳
	rand.Seed(time.Now().UnixNano())

	// 这里将 len(intArr) 函数赋值给 len 变量，下面直接引用 len 变量即可，因为在程序中调用 len() 函数特别耗费内存资源
	len := len(intArr)

	for i := 0 ; i < len ; i++ {
		// 随机数就是大于等于0 小于一百之间，并且赋值给 intArr 数组
		intArr[i] = rand.Intn(100) 
	}
	// 交换前打印
	fmt.Printf("intArr随机数 = %v\n",intArr)

	// 定义一个临时的变量，用于交换
	temp := 0

	// 交换的次数是数组的长度除以 2，也就是交换次数的循环次数
	for i := 0 ; i < len / 2 ; i++ {

		// 先将倒数第一个元素赋值给 temp，这里是 len - 1 , 因为不减一的话就会越界
		// 为什么要减 i：因为在循环第二次的时候此时 i 就变 1 了然后再一减刚好是倒数第二个
		// temp = intArr[4]
		temp = intArr[len -1 -i]

		// 再将正数的第 n 个元素，交换给倒数的第 n 个元素
		// intArr[4] = intArr
		intArr[len -1 -i] = intArr[i]

		// 这里是将 temp 变量中的倒数第一个值交给了 intArr[i] 这个对应的正数元素
		// intArr[0] = temp(intArr[4])
		intArr[i] = temp 

		// 然后依次 for 循环将 i 进行改变依次类推进行赋值
	}
	// 交换后打印
	fmt.Printf("交换后 = %v",intArr)
}