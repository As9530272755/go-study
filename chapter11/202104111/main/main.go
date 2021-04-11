package main

import (
	"fmt"
	"go_code/exec/202104111/model"
)

func main()  {
	a := model.NewAccount("")
	a.SetAccoutn("zgy123")
	a.SetBalance(30)
	a.SetPwd("666666")
	fmt.Println("账户：",a.GetAccoutn(),"\n余额：",a.GetBalance(),"\n密码：",a.GetPwd())
}