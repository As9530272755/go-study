package main
import	(
	"fmt"
)

func main()  {
	// 演示关系运算符的使用
	var n1 int = 9
	var n2 int = 8

	fmt.Println(n1 == n2) // n1=9 n2=8 这是判断 9 等于 8 吗，返回 false
	fmt.Println(n1 != n2) // n1=9 n2=8  9 不等于 8 返回 true
	fmt.Println(n1 > n2) // 9 大于 8 返回 true
	fmt.Println(n1 >= n2) // 9 大于或者等于 8 返回 true
	fmt.Println(n1 < n2) // 9 小于 8 返回 false
	fmt.Println(n1 <= n2) // 9 小于或等于 8 返回 false

	// 通过变量赋值的方式执行，我们将 n1 > n2 的通过类型推导赋值给 flag 变量
	flag := n1 > n2
	fmt.Println("flag=",flag)
}