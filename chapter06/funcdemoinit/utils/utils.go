package utils

// Age 和 Name 全局变量，我们需要在 main.go 使用
// 但是我们需要初始化 Age 和 Name
var Age int
var Name string

// 让 init 函数完成初始化工作
func init()  {
	Age  = 100
	Name = "Tom"
}