package main
import (
	"fmt"
)

//演示golang中基本数据类型的转换
func main(){

	var i int32 = 100
	//希望将 i 转成一个float类型 
	var n1 float32 = float32(i)
	var n2 int8 = int8(i)//这里是先将i转化为int8类型交给了n2
	var n3 int64 = int64(i)//int64也是需要转的 低精度 -> 高精度
	fmt.Printf("i=%v n1=%v n2=%v n3=%v \n",i,n1,n2,n3)

	//被转换的是变量存储的数据(即值)，变量本身的数据类型并没有变化！
	fmt.Printf("i type is %T",i)

	//只是转换的结果是按溢出处理溢出处理，和我们希望的结果不一样。 int8【-128---127】 
	var num1 int64 = 999999  //这个值在int64中可以存放
	var num2 int8 = int8(num1) //这里我们定义一个num2 数据类型为int8 而且是转换了数据类型
	fmt.Println("num2=",num2)	//这时候我们保存之后编译器并没有报错
}
