package main
import (
	"fmt"
	"unsafe"
)

//演示golang中布尔类型的使用
func main(){
	var b = false
	fmt.Println("b=", b)
	//在使用的时候需要注意bool类型占用的存储空间是一个字节

	fmt.Println("b占用的空间=" ,unsafe.Sizeof(b) )
}