package main

import (
	"fmt"
)

// 定义 Account 结构体
type Account struct {
	User string
	Pwd string
	Blance float64
}

// 定义方法
// 第一个方法可以存款
func (account *Account) Deposite() {
	
	// 提示当前功能
	fmt.Println("当前功能：存款！")
	
	var user string
	fmt.Println("请输入用户名:")
	fmt.Scanln(&user)

	var pwd string
	fmt.Println("请输入密码:")
	fmt.Scanln(&pwd)

	var money float64
	fmt.Println("请输入存款金额:")
	fmt.Scanln(&money)

	// 既然要存款先要匹配密码用户名
	// 如果我们输入的密码和 Account 结构体中的密码不匹配就提示输出的密码不正确
	if pwd == account.Pwd && user == account.User {
		
		// 检测存放的钱是否正确，如果存放的钱都不大于零就是误操作
		if money <= 0 {
			fmt.Println("输入的存款金额不正确！")
			return
		}
	} else {
		fmt.Println("账户名或密码输入不正确！")
		return
	}
	
	// 如果上面都没有错误，就讲存放的金额在 Account 结构体 Blance 字段的原有基础上增加
	account.Blance += money
	fmt.Println("存款成功，现有余额：",account.Blance)
}

// 第二个方法可以取款，因为取款需要输入存款金额和密码所以要使用 money 和 pwd 这两个形参
func (account *Account) WithDraw() {
	// 提示当前功能
	fmt.Println("当前功能：取款！")

	var user string
	fmt.Println("请输入用户名:")
	fmt.Scanln(&user)

	var pwd string
	fmt.Println("请输入密码:")
	fmt.Scanln(&pwd)

	var money float64
	fmt.Println("请输入存款金额:")
	fmt.Scanln(&money)

	// 先匹配密码账户名是否正确，如果用户密码正确就判断取款金额是否正确
	if user == account.User && pwd == account.Pwd {
		
		// 取款金额是否输入正确,如果取得钱小于 0 或者大于了当前余额就说明取款金额不对
		if money <= 0 || money > account.Blance {
			fmt.Println("输入的存款金额不正确！")
			return
		}
	} else {
		fmt.Println("账户名或密码输入不正确！")
		return
	}

	account.Blance -= money
	fmt.Println("取款成功，现有余额：",account.Blance)
}

// 定义第三个方法，查询余额。因为查询余额需要登录所以需要使用到 user 和 pwd 两个形参
func (account *Account) Query() {
	// 提示当前功能
	fmt.Println("当前功能：查询余额！")
	
	var user string
	fmt.Println("请输入用户名:")
	fmt.Scanln(&user)

	var pwd string
	fmt.Println("请输入密码:")
	fmt.Scanln(&pwd)

	// 判断我们输入的用户还有密码都正确的话就是输出当前余额
	if user == account.User && pwd == account.Pwd {
		fmt.Println("当前余额：",account.Blance)
	} else {
		fmt.Println("账户名或密码输入不正确！")
		return
	}
}

func main() {
	// 测试，先给 Account 结构体赋初始值,这里定义变量 account 类型为 Account 结构体
	account := Account {
		User : "zgy",
		Pwd : "666666",
		Blance : 10.0,
	}

	// 定义 open 变量然后通过键盘输入想要执行的功能
	var open int
	fmt.Println("请输入想要操作的功能：\n1.查询余额\n2.存款\n3.取款")
	fmt.Scanln(&open)
	switch open {
		case 1:
			account.Query()
		case 2:
			account.Deposite()
		case 3:
			account.WithDraw()
		default:
			fmt.Println("请输入正确的操作")
	}

}