package main

import (
	"fmt"
)

// 定义了 sum 函数，形参是 int 类型，返回值也是 int 
func sum(n1 int,n2 int) int {
	
	//当我们的代码执行到 defer 的时候，我们的编译器会将 defer fmt.Println("ok1 n1=", n1) 压入到一个栈中，也就是说执行到 defer 的时候暂时不执行 fmt.Println("ok1 n1=", n1) 这句话，而是讲这句话压入到独立的(defer)的栈中。这个栈是一个独立的栈和我们的 sum、main 栈不在同一个地方
	defer fmt.Println("ok1 n1=", n1) // 第四次执行该 defer 语句：遵循先入后出的原则
	
	// 下面这个 defer 也是同理
	// 什么时候执行 defer 呢：当函数执行完毕后再从这个 defer 栈中按照先入后出的方式出栈然后执行（也就是说谁第一个 defer 语句进入反而最后执行，类似于堆盘子取盘子的方法）
	defer fmt.Println("ok2 n2=", n2) // 第三次执行 defer ：遵循先后出的原则
	
	
	res := n1 + n2 // res = 30
	fmt.Println("ok3 res=",res) // 第一次执行： ok3 res= 30
	return res					// 第二次执行： return res ，这时候就表示 sum 函数执行完毕了，当函数执行完毕之后就开始在 defer 栈中执行东西了。
}

func main()  {
	// 主函数调用了上面自定义的 sum 函数，将 10 和 20 分别传给 n1，n2
	res := sum(10,20) 
	fmt.Println("main()res=",res) // 第五次执行：因为在程序中都是从上往下执行，当上一个 sum 函数执行完毕之后在执行 main 函数
}

// 该程序执行顺序是：
// 1.fmt.Println("ok3 res=",res) 
// 2.return res
// 3.defer fmt.Println("ok2 n2=", n2)
// 4.defer fmt.Println("ok1 n1=", n1)
// 5.fmt.Println("main()res=",res) 