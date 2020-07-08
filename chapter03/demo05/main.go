package main
import "fmt"

//演示golang中+号的使用
func main(){

	var str1 = "hello"		// str1 是一个变量、然后赋值为hello
	var str2 = "world"		// str2 是一个变量、然后赋值为world
	var str3 = str1 + str2	// str3 这个变量的结果是通过 str1 + str2 这两个变量的值相加
	fmt.Println("str3=",str3)
}