package main
import (
	"fmt"
)
//演示golang中string（字符串）类型使用
func main(){

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

	var str = "hello world"
	str += "haha"
	fmt.Println(str)

	var str2 = "hello" + "world" + "hello" + "world" + 	// 加号一定要放到当前代码的末尾
	"hello" + "world" + "hello" + "world" + 		// 加号一定要放到当前代码的末尾
	"hello" + "world" + "hello" + "world" + 		// 加号一定要放到当前代码的末尾
	"hello" + "world"								// 结尾不能再加上 +
	fmt.Println(str2)

	var a int 
	var b float32
	var c float64
	var isMarried bool
	var name string
	fmt.Println(a,b,c,isMarried,name)
}