package main

import (
	"fmt"
)

// 定义一个 modefyUser 的函数
func modifyUser(users map[string]map[string]string,name string)  {
	
	// 3.modifyuser 函数拿到了 main 函数中传递过来的 users map 之后进判断是否有这个 name key
	// 如果 users name 中有这个用户了就不会等于 nil，如果没有就会等于 nil
	if users[name] != nil {
		
		// 有这个用户了,将其 密码 给为 888888
		users[name]["pwd"] = "888888"
	} else {
		
		// 没有这个用户就要增加这个用户信息，首先就需要将其 make 在内存中开辟一个空间来存放
		users[name] = make(map[string]string)
		users[name]["pwd"] = "888888"
		users[name]["nickname"] = "昵称" + name
	}
}

func main()  {
	// 1.在主函数中定义一个 map
	users := make(map[string]map[string]string)
	// 测试，这里放入 zgy 这个用户，他的值对应的是 map[string]string
	users["zgy"] = make(map[string]string)
	users["zgy"]["pwd"] = "123456"
	
	// 2.将主函数中的 users map 传递给 modifyUser 函数中，tom 传递给 modifyUser 函数中的 name
	modifyUser(users,"tom")
	modifyUser(users,"zgy")
	
	// 最后输出
	fmt.Println(users)
}