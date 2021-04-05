package main

import (
	"fmt"
)

type Calcuator struct {
	Num1 float64
	Num2 float64
}

/*
定义小小计算器结构体(Calcuator)，实现加减乘除四个功能
实现形式1: 分四个方法完成
实现形式2: 用一个方法搞定
*/

func (ma *Calcuator) getSum() float64 {
	return ma.Num1 + ma.Num2
}

func (ma *Calcuator) getSum1() float64 {
	return ma.Num1 - ma.Num2
}

func (ma *Calcuator) getSum2() float64 {
	return ma.Num1 * ma.Num2
}

func (ma *Calcuator) getSum3() float64 {
	return ma.Num1 / ma.Num2
}

func main()  {
	var mu Calcuator
	mu.Num1 = 10.0
	mu.Num2 = 2.0
	res := mu.getSum1()
	fmt.Println("计算结果:",res)
}