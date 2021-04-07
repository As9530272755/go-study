package main

import (
	"fmt"
)

type Ponts struct {
	x int
	y int
}

type Btyr struct {
	x int
	y int
}

func main()  {
	var a Ponts
	var b Btyr

	a = Ponts(b)
	a.x = 1
	a.y = 2
	fmt.Println(a,b)
}
