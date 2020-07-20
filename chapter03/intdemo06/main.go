package main
import (
	"fmt"
	"unsafe"
)

func main(){

	//整型的使用细节
	var n1 = 100 //请问n1是什么类型
	//这里我们给大家介绍如何查看某个变量的数据类型
	fmt.Printf("n1 的 类型 %T",n1) //这里的fmt.Print可以用来做格式化输出、说白了就是可以用来输出东西

	//如何在程序查看某个变量的字节大小和数据类型
	var n2 int64 = 10 //我们现在看一下n2他的字节大小现在是多少和数据类型又是什么
	
	//unsafe.Sizeof(n2)是unsafe 包的一个函数，可以返回n2变量的占用字段数
	fmt.Printf("n2 的 类型 %T n2占用的字节数是 %d\n",n2,unsafe.Sizeof(n2))
	//上面代码意思Printf是用于格式化的%T就是n2的这个类型%d就是相当于将unsafe.Sizeof(n2)的返回值进行打印


	//Golang 程序中整型变量在使用时，遵守保小不保大的原则
	//即：在保证程序正确运行下，尽量使用占用空间小的数据类型。【如：年龄】
	var age	byte = 90	//在此我们使用byte整数类型、因为byte的取值范围是0-255
	fmt.Println("age=",age)
}