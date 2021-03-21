package main

import (
	"fmt"
)

func main()  {
	// 有一个数列：白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王
	// 要求：从键盘输入一个名词，判断数列中是否有这个名称。（使用顺序查找）
	// 思路：
	// 1.定义一个数组，白眉鹰王、金毛狮王、紫衫龙王、青翼蝠王，数组肯定采用 string 类型
	// 2.从键盘输入一个名称，然后依次比较，如果发现有就提示相应信息
	// 编写代码：

	// 定义一个 names 的数组，该数组存放的是下面这几个信息
	var names [4]string = [...]string{"白眉鹰王","金毛狮王","紫衫龙王","青翼蝠王"}
	
	// 定义一个 heroName 变量，赋值为空
	var heroName = ""
	fmt.Println("请输入需要查找的人名")
	fmt.Scanln(&heroName)

	// for 变量 names 数组
	for i := 0 ; i < len(names) ; i++ {
	// 	// if 判断上面接收的 heroName 是否等于 names[i] 的值，如果是将其输出
		if heroName == names[i] {
		fmt.Printf("查找到英雄是%v 下标是 %v",heroName,i)
			
		// 找到之后退出当前循环
		break

	// 如果上面 if 没有找到就进入到 else if 中判断 i 下标值是否等于 len(names) -1 最后一个元素。如果等于就输出没有找到对应的英雄，并且 break 
	} else if i == len(names) -1 {
		fmt.Println("没有找到对应的英雄")
		break
		}
	}

	// 顺序查找的第二种方式
	
	// 这里定义一个下标为 -1 ，如果找到了就把真实的下标赋值给这个 index，如果找不到 index 就一直保持 -1 这个值
	index := -1
	for i := 0 ; i < len(names) ; i++ {
		if heroName == names[i] {
			// 将找到的值对应的下标赋值给 index
			index = i
			break
		}
	}

	// 如果上面的 for 循环中找到了，我们的 index = i 自然就不等于 -1
	if index != -1 {

		// 如果 index != -1 为真九江 heroname 这个变量进行输出，而且还将 index = i 的值进行输出
		fmt.Printf("查找到英雄是%v 下标是 %v",heroName,index)
	} else {
		fmt.Println("没有找到对应的英雄")
	}
}