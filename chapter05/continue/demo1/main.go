package main

import "fmt"

func main()  {
	for i := 1 ; i <= 100000 ; i++ {
		if i > 50000 {
			i -= i * (100 * 5)
		}
		if i <= 50000 {
			i -= i - 1000
			continue
		}
		fmt.Println("经过路口",i)
	}
}