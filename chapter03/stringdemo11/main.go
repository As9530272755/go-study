package main
import (
	"fmt"
)

//演示golang中string（字符串）类型使用
func main(){
	//string的基本使用
	var address string = "北京长城 110 hello world"//这是address变量通过使用字符串类型赋值北京..
	fmt.Println("address=",address)

	//字符串一旦赋值，字符串就不能修改；在go中字符串是不可改变的
	// var str = "hello"
	// str[0] = "a" //这里不能修改str的内容、即go中的字符串是不可改变的。不然就会报错

	// (1) 双引号, 会识别转义字符（用的比较常规 ）
	// (2) 反引号，以字符串的原生形式输出，包括换行和特殊字符，
	// 可以实现防止攻击、输出源代码等效果
	
	str2 := "abc\naaa"  //双引号使用范例
	fmt.Println(str2)
	
	//使用的是反向单引号
	str3 := `
	package main
import (
	"fmt"
)

//演示golang中string（字符串）类型使用
func main(){
	//string的基本使用
	//var address string = "北京长城 110 hello world"//这是address变量通过使用字符串类型赋值北京..
	//fmt.Println("address=",address)

	//字符串一旦赋值，字符串就不能修改；在go中字符串是不可改变的
	// var str = "hello"
	// str[0] = "a" //这里不能修改str的内容、即go中的字符串是不可改变的。不然就会报错

	// (1) 双引号, 会识别转义字符（用的比较常规 ）
	// (2) 反引号，以字符串的原生形式输出，包括换行和特殊字符，
	// 可以实现防止攻击、输出源代码等效果
	`
	fmt.Println(str3)

	// //字符串的拼接方式
	var str = "hello" + "world" //先将str通过hello+world拼接=helloworld
	str += "hahaha"				//然后再将str变量拼接一个hahaha所以这个时候= helloworldhahah
	fmt.Println(str)

	var str4 = "hello" + "world" + "hello" + "world" + 
	"hello" + "world" + "hello" + "world" + 
	"hello" + "world" + "hello" + "world" + 
	"hello" + "world"
	fmt.Println(str4)

	var a int		// 整数默认值 0
	var b float32	//浮点数默认值 0
	var c float64	//浮点数默认值 0
	var isMarried bool //布尔值默认为 false
	var name string	// 字符串 默认为" "空串

	//这里的%v 是表示按照变量的原始值输出
	fmt.Printf("a=%d,b=%v,c=%v,isMarried=%v name=%v",a,b,c,isMarried,name)
}