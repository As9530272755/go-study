package Model

import (
	"fmt"
)

// 定义 customer 结构体，表示一个客户信息
type Customer struct {
	Id int
	Name string
	Gender string
	Age int
	Phone string
	Status int
}

// 提供工厂函数，返回一个 Customer 的实例
func NewCustomer(id int,name string,gender string,age int,phone string,status int) Customer {
	return Customer {
		Id : id,
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Status : status,
	}
}

// 第二种创建 customer 方法，不带 id 的工厂函数
func NewCustomer2(name string,gender string,age int,phone string,status int) Customer {
	return Customer {
		Name : name,
		Gender : gender,
		Age : age,
		Phone : phone,
		Status : status,
	}
}

// 返回用户信息，格式化的字符串
func (this Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t%v",
	this.Id,this.Name,this.Gender,this.Age,this.Phone,this.Status)
	return info
}