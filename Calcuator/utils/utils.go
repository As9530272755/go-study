package utils

import (
	"fmt"
)

type Calcuator struct {
	Num1 float64
	Num2 float64
	Exit bool
}

func NewCalcuator() *Calcuator {
	return &Calcuator {
		Num1 : 0.0,
		Num2 : 0.0,
	}
}

func (calcuator *Calcuator) sum() {
	fmt.Println("这是加法计算")
	fmt.Println("请输入第一个数😘：")
	fmt.Scanln(&calcuator.Num1)
	fmt.Println("请输入第二个数🤳：")
	fmt.Scanln(&calcuator.Num2)
	res := calcuator.Num1 + calcuator.Num2
	fmt.Println("最终结果🤣",res)
}

func (calcuator *Calcuator) sub() {
	fmt.Println("这是减法计算")
	var res float64
	fmt.Println("请输入第一个数😘：")
	fmt.Scanln(&calcuator.Num1)
	fmt.Println("请输入第二个数🤳：")
	fmt.Scanln(&calcuator.Num2)
	if calcuator.Num1 < calcuator.Num2 {
		fmt.Println("请输入正确的数值🈲")
		return
	} else {
		res = calcuator.Num1 - calcuator.Num2
		fmt.Println("最终结果🤣",res)
	}
}

func (calcuator *Calcuator) ride() {
	fmt.Println("这是乘法计算")
	var res float64
	fmt.Println("请输入第一个数😘：")
	fmt.Scanln(&calcuator.Num1)
	fmt.Println("请输入第二个数🤳：")
	fmt.Scanln(&calcuator.Num2)
	if calcuator.Num1 == 0 || calcuator.Num2 == 0 {
		fmt.Println("输入的数不能为 0 哟😁")
		return
	} else {
		res = calcuator.Num1 * calcuator.Num2
		fmt.Println("最终结果🤣",res)
	}
}

func (calcuator *Calcuator) division() {
	fmt.Println("这是除法计算")
	var res float64
	fmt.Println("请输入第一个数😘：")
	fmt.Scanln(&calcuator.Num1)
	fmt.Println("请输入第二个数🤳：")
	fmt.Scanln(&calcuator.Num2)
	if calcuator.Num1 == 0 || calcuator.Num2 == 0 {
		fmt.Println("输入的数不能为 0 哟😁")
	} else {
		res = calcuator.Num1 / calcuator.Num2}
	fmt.Println("最终结果🤣",res)
}

func (calcuator *Calcuator) exit() {
	for {
		var exit string
		fmt.Printf("你确定想退出吗😭\n我可舍不得你呀\n退出请按 y \n不想退出请按 n\n")
		fmt.Scanln(&exit)
		if exit == "y" {
			calcuator.Exit = true
			break
		} else if exit == "n" {
			return
		}
	}
}

func (calcuator *Calcuator) Menu() {
	for {
		fmt.Println("这是个🐱‍💻🐱‍🚀🐱‍👓🐱‍👤计算器")
		fmt.Println()
		var open string
		fmt.Printf("请输入需要计算的(+-*/)\n退出请按4：\n")
		fmt.Scanln(&open)
		switch open {
			case "+":
				calcuator.sum()
			case "-":
				calcuator.sub()
			case "*":
				calcuator.ride()
			case "/":
				calcuator.division()
			case "4":
				calcuator.exit()
			default:
				fmt.Println("请输入正确的字符")
		}
		if calcuator.Exit != false {
			fmt.Println("退出🐱‍💻🐱‍🚀🐱‍👓🐱‍👤计算器🥶")
			return
		}
	}
}