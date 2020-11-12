package main

import _ "fmt" // 如果我们没有使用到一个包，但是又想去掉，前面可以加一个 _ 表示忽略

func main()  {
	// var n1 int32 = 12
	// var n2 int64
	// var n3 int8

	// n2 = int64(n1) + 20		// 通过强转换的凡是将 n1 的 int32 转换为 n2  int64
	// n3 = int8(n1) + 20		// 通过强转换的凡是将 n1 的 int32 转换为 n3  int 8
	// fmt.Println(n1,n2,n3)	// 在输出的时候通过编译


// 		var n1 int32 = 12	
// 		var n3 int32
// 		var n4 int32
// 		n4 = int32(n1) + 127 // 编译能够通过但是最终的结果是会溢出，因为 int8 数据类型最大值为 127 。就会按照溢出给一个随机结果
// 		n3 = int32(n1) + 128 // int8 数据结构最大只能够到 -128~127 ，所以当编译器一看 128 已经超了 int8 数据结构的最大值直接就编译不过
// 		fmt.Println(n3,n4)		// 
}