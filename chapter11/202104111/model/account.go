package model

import (
	"fmt"
)

// 编写 account1 结构体，其有三个私有字段
type account struct {
	account1 string
	balance float64
	pwd string
}

// 定义一个函数，返回值为 account 结构体
func NewAccount(account1 string) *account {
	return &account {
		account1 : account1,
	}
}

// 定义 SetAccoutn 方法
func (a *account) SetAccoutn(account string) {
	if len(account) >= 6 && len(account) <= 10 {
		a.account1 = account
	} else {
		fmt.Println("请输入正确的账户长度（6-10）：")
	}
}

// 定义 GetAccoutn 方法
func (a *account) GetAccoutn() string {
	return a.account1
}

// 定义 SetBalance 方法
func (a *account) SetBalance(balance float64) {
	if balance > 20 {
		a.balance = balance
	} else {
		fmt.Println("请输入正确的余额：")
	}
}

// 定义 GetBalnce 方法
func (a *account) GetBalance() float64 {
	return a.balance
}

// 定义 SetPwd 方法
func (a *account) SetPwd(pwd string) {
	if len(pwd) == 6 {
		a.pwd = pwd
	} else {
		fmt.Println("请输入正确的密码：")
	}
}

// 定义 GetPwd 方法
func (a *account) GetPwd() string {
	return a.pwd
}