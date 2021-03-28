package main

import (
	"fmt"
)

// 定义一个学生的结构体
// 这个 stu 的结构体中存放了三个属性，分别是 name、age、address
type Stu struct {
	Name string
	Age int
	Address string
}

func main()  {
	// 1.map 的 key 为学生的编号，并且这个学号是唯一的
	// 2.map 的 value 为结构体，包含学生的 名字、年龄、还有地址

	// 3.开始声明 map ,这个 map 的 key 还是 string 但是他的 value 是 stu 结构体
	students := make(map[string]Stu)

	// 创建两个学生,这两个学生的值调用 Stu 结构体，并赋值给 stu1 和 stu2 这两个变量
	stu1 := Stu{"张三",25,"重庆"}
	stu2 := Stu{"黄二狗",30,"南昌"}

	// 现在需要将 stu1 和 stu2 交给 students 这个 map
	// students["no1"] 为 key stu1 就是 value
	students["no1"] = stu1
	students["no2"] = stu2

	fmt.Println(students)
	fmt.Println()

	// 遍历各个学生的信息
	// 因为他已经是 map 类型了，直接可以通过 for-range 方式遍历即可
	for k,v := range students {

		// 这里的 k 就是 students map 中的 key
		fmt.Printf("学生的编号是 %v \n",k)

		// 由于 students map 的 value 是 Stu 结构体，所以用 v.跟结构体中的属性即可取出
		fmt.Printf("学生的姓名是是 %v \n",v.Name)
		fmt.Printf("学生的年龄是 %v \n",v.Age)
		fmt.Printf("学生的地址是 %v \n",v.Address)
		fmt.Println()
	}
}