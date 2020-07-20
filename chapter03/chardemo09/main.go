package main
import (
	"fmt"
)

//演示golang中字符类型得使用
func main(){

	//var c1 byte = 'a'  //这是保存a并且是通过byte来保存
	//var c2 byte = '0'	//这是保存字符的0
	
	//当我们直接输出byte值时，就时输出了他对应的字符的ascll码值
	//因为小写的a对应的ascll码十进制是97 
	//在看对应的0他对应的ascll码是对应的十进制48
	//fmt.Println("c1=",c1)		//输出c1的变量
	//fmt.Println("c2=",c2)		//输出c2的变量

	//如果我们希望输出的是对应的字符、需要使用格式化输出
	//fmt.Printf("c1=%c c2=%c", c1,c2 )	//这样就是格式化输出、告诉我们的编译器c1就输出c1的值、c2就输出c2的值

	//var c3 int = '北'  	//overflows 溢出

	//var c4 int = 22269 //22269对应的就是unicode码中的中国的数值

	//字符类型是可以进行运算的、相当于一个整数、运算时是按照码值运算
	var n1 =10 + 'a' //a 等于UTF-8码中的97，所以的到的结果九尾10+97=107
	fmt.Println("n1=",n1 )
}