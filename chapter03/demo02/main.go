package main
import "fmt"

func main(){
	//第三种：省略 var
	//有个name的变量但是我们没有使用关键字、我们如果不使用关键字就要通过:的方式来进行声明赋值
	//下面的方式等价 var name string  name = "zhang"
	// := 这个 : 不能省略，否则就是错的
	name := "zhang"
	fmt.Println("name=",name)
}