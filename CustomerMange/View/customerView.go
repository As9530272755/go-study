package main

import (
	"fmt"
	"go_code/CustomerMange/Service"
	"go_code/CustomerMange/Model"
)

type customerView struct {
	// 定义必要的字段
	// key 接收用户的输入
	key string

	// exit 表示是否退出该程序
	exit bool

	// 增加一个字段 customrService 字段，调用 Service 包中的 CustomerService 结构体
	customerservice *Service.CustomerService
	customerslice Model.Customer
}

// 显示所有的客户信息
func (this *customerView) list() {
	
	// 首先，获取到当前所有的客户信息都在(切片中),调用 customerservice.List() 方法，该方法返回的是一个 []Model.Customer 切片
	customers := this.customerservice.List()

	// 显示客户信息，制作表头和表尾
	fmt.Println("--------------------------客户列表--------------------------------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t身份证")
	
	// 遍历 service 中的所有信息
	for i := 0 ; i < len(customers) ; i++ {
		fmt.Println(customers[i].GetInfo())
	}
	fmt.Println("--------------------------客户列表完成----------------------------")
	fmt.Println()
	fmt.Println()
}

// 得到用户的输入信息，构建新的客户，并完成添加
func (this *customerView) Add() {
	fmt.Println("--------------------------添加客户--------------------------------")
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("身份证：")
	status := 0
	fmt.Scanln(&status)

	// 构建新的 customer 实例
	// 注意：id 号没有让用户输入，因为 id 是唯一的，需要系统分配
	// 定义 customer 变量，赋值时 Model.NewCustomer2 函数 ，并且传入 name,gender,age,phone,status 信息
	customer := Model.NewCustomer2(name,gender,age,phone,status)
	
	// 调用 customerservice 中的 Add 方法，并且传入 customer 变量中的值
	if 	this.customerservice.Add(customer) {
		fmt.Println("--------------------------添加完成--------------------------------")
	} else {
		fmt.Println("--------------------------添加失败--------------------------------")
	}

}

// 通过用户输入，删除该 id 
func (this *customerView) Delete()  {
	fmt.Println("--------------------------删除客户--------------------------------")
	fmt.Println("请选择待删除的客户编号(输入 -1 为退出删除)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃删除操作
	}

	fmt.Println("确认是否删除(Y/N)：")
	choice := ""
	fmt.Scanln(&choice)
	
	if choice == "Y" || choice == "y" {
		// 调用 service 中 delete 方法，并传入 id
		if this.customerservice.Delete(id) {
			fmt.Println("--------------------------删除成功--------------------------------")
		} else {
			fmt.Println("\n--------------删除失败，输出的 id 号不存在，请重新输入---------------")
		}
	}
}

func (this *customerView) Exit() {
	for {
		exit := ""
		fmt.Println("确认是否退出(Y/N)")
		fmt.Scanln(&exit)
		if exit == "y" || exit == "Y" {
			this.exit = false
			return
		} else if exit == "n" || exit == "N" {
			this.exit = true
			return
		}
		if exit != "Y" && exit != "y" {
			fmt.Println("输入有误！请重新输入(Y/N)")
		} else if exit != "n" && exit != "N" {
			fmt.Println("输入有误！请重新输入(Y/N)")
		}
	}
}

func (this *customerView) Revise() {
	fmt.Println("--------------------------修改客户--------------------------------")
	fmt.Println("请选择待修改的客户编号(输入 -1 为退出修改)：")
	id := -1
	fmt.Scanln(&id)
	if id == -1 {
		return // 放弃修改操作
	}
	
	fmt.Println("姓名：")
	name := ""
	fmt.Scanln(&name)
	fmt.Println("性别：")
	gender := ""
	fmt.Scanln(&gender)
	fmt.Println("年龄：")
	age := 0
	fmt.Scanln(&age)
	fmt.Println("电话：")
	phone := ""
	fmt.Scanln(&phone)
	fmt.Println("身份证：")
	status := 0
	fmt.Scanln(&status)

	// 构建新的 customer 实例
	// 注意：id 号没有让用户输入，因为 id 是唯一的，需要系统分配
	// 定义 customer 变量，赋值时 Model.NewCustomer2 函数 ，并且传入 name,gender,age,phone,status 信息
	customer := Model.NewCustomer2(name,gender,age,phone,status)
	
	// 调用 customerservice 中的 Add 方法，并且传入 customer 变量中的值
	if 	this.customerservice.Revise(id,customer) {
		fmt.Println("--------------------------修改完成--------------------------------")
	} else {
		fmt.Println("\n--------------------------修改失败--------------------------------")
	}

}

// 显示主菜单
func (cv *customerView) mainMenu() {
	for {
		fmt.Println("--------------------------客户信息管理软件------------------------")
		fmt.Println("                      1、添加客户                       ")
		fmt.Println("                      2、修改客户                       ")
		fmt.Println("                      3、删除客户                       ")
		fmt.Println("                      4、客户列表                       ")
		fmt.Println("                      5、退    出                       ")
		fmt.Println()
		fmt.Println("                      请选择(1-5):                      ")

		fmt.Scanln(&cv.key)
		switch cv.key {
			case "1":
				cv.Add()
			case "2":
				cv.Revise()
			case "3":
				cv.Delete()
			case "4":
				cv.list()
			case "5":
				cv.Exit()
				break
			default :
				fmt.Println("输入有误！请重新输入！")
		}
		if !cv.exit {
			fmt.Println("已经退出客户信息管理软件!")
			return
		}
	}
}

func main()  {
	// 创建 customer 实例调用 customer.mainMenu() 方法
	// key 默认给 空字符串， exit 默认为 true 反复 for 循环
	customer := customerView{
		key : "",
		exit : true,
	}

	// 这里完成对 customerView 结构体的 customerservice 字段初始化工作
	customer.customerservice = Service.NewCustomerService()

	// 调用主菜单
	customer.mainMenu()
}