package utils

import (
	"fmt"
)

type myAccount struct {

	// 声明必要的字段
	name string
	pwd int
	key string
	loop bool
	balance float64
	money float64
	note string
	flag bool
	details string
}


// 编写工厂模式，返回 myAccount 结构体实例,并且定义定义初始值
func NewmyAccount() *myAccount {
	return &myAccount{
		name : "张桂元",
		pwd : 123456,
		key : "",
		loop : true,
		balance : 0.0,
		money : 0.0,
		note : "",
		flag : false,
		details : "支出-----账户金额-----收支金额-----说明",
	}
}

// 给 myAccount 结构体绑定相应方法
// 添加登录方法
func (my *myAccount) login() {
	fmt.Println("欢迎使用老张记账软件！！！")
	name := ""
	tot := 3
	var pwd int
	for i := 0 ; i < 3 ; i++ {
		fmt.Println("请输入用户名：")
		fmt.Scanln(&name)
		fmt.Println("请输入密码：")
		fmt.Scanln(&pwd)
		if name == my.name && pwd == my.pwd {
			fmt.Println("登录成功！！！")
			break
		} else {
			tot--
			if tot != 0 {
			fmt.Printf("请输入正确的用户名或密码！\n当前还有 %v 次机会请再次输入！\n",tot)
			}
		}

		// 如果 i = 2 时，my.loop = false 退出程序
		if i == 2 {
			my.loop = false	
		}
	}
}

// 将显示明细绑定一个方法
func (my *myAccount) showDetails() {
	fmt.Println("------------当前收支明细记录-----------")
	// 如果有收支 flag 就等于 true 输出 details，否则提示来一笔收支
	if my.flag == true {
		fmt.Println(my.details)
	} else {
		fmt.Println("当前没有收支明细...来一笔吧!" )
	}
}

// 将登记收入写成一个方法
func (my *myAccount) income() {
	fmt.Println("------------当前登录收入记录-----------")
	fmt.Println("本次收入金额：")
	fmt.Scanln(&my.money)
	my.balance += my.money
	fmt.Println("本次收入说明：")
	fmt.Scanln(&my.note)
	my.details += fmt.Sprintf("\n收入     %v        %v        %v",my.balance,my.money,my.note)
	
	// 如果有收入就将 flag 变为 true
	my.flag = true
}

// 将支出绑定为一个方法
func (my *myAccount) expend() {
	
	fmt.Println("------------登记支出-----------")
	fmt.Println("本次支出金额：")
	fmt.Scanln(&my.money)
	
	if my.money > my.balance {
		 fmt.Println("支出余额不足，不能支出")
	} else if my.money < my.balance {
			my.balance -= my.money
	}
		
	fmt.Println("本次支出说明：")
	fmt.Scanln(&my.note)

	my.details += fmt.Sprintf("\n支出     %v        %v        %v",my.balance,my.money,my.note)
	// 如果有支出就将 flag 变为 true
	my.flag = true
}

// 绑定退出软件的方法
func (my *myAccount) exit() {
	exit := ""
	fmt.Println("确定要退出吗?（y/n）")
	for {
		fmt.Scanln(&exit)
		if exit == "y"  {
		my.loop = false	
		break	
		} else if exit == "n" {
			break
		}
		fmt.Println("输入有误，请输入 y 或 n ！！！")
	}
}

// 绑定一个方法显示主菜单
func (my *myAccount) MainMenu() {
	my.login()	
	if !my.loop {
		fmt.Println("由于账户名或者密码不正确退出程序！")
		return
	}
	for {
			fmt.Println("\n------------家庭收支记账软件-----------")
			fmt.Println("                1 收支明细")
			fmt.Println("                2 登记收入")
			fmt.Println("                3 登记支出")
			fmt.Println("                4 退    出")
			fmt.Println()

			fmt.Println("请选择(1-4)：")
			fmt.Scanln(&my.key)
			switch my.key {
				case "1":
					my.showDetails()
				case "2":
					my.income()
				case "3":
					my.expend()
				case "4":
					my.exit()
				default:
					fmt.Println("请输入正确的选项")
			}
		if !my.loop {
			break
		}
	}
	fmt.Println("您已经退出了该软件！")
}