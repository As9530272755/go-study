package main
import "fmt"

//下面这三个var是定义全局变量
var n1 = 100
var n2 = 200
var name = "jack"

//下面这种方式和上面这种凡是也是一次性声明、并且两者之间是等价的
var (
	n3 = 77
	n4 = 99
	name2 = "gui"
)

func main(){

	//输出上面定义的三个全局变量
	fmt.Println("n1=",n1 ,"n2=",n2 ,"name=",name)
	fmt.Println("n3=",n3 ,"n4=",n4 ,"name2=",name2)
} 