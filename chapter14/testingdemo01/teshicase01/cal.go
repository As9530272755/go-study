package main

func AddUpper(n int) int {
	res := 0
	for i := 0; i <= n; i++ {
		res += i
	}
	return res
}

func Sub(n1 int, n2 int) int {
	return n1 - n2
}
