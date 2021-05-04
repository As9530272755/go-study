package main

import (
	"flag"
	"fmt"
)

func main() {
	var (
		user string
		pwd  int
		host string
		port int
	)

	flag.StringVar(&user, "u", "", "user 默认为空")
	flag.IntVar(&pwd, "p", 0, "pwd 默认为0")
	flag.StringVar(&host, "h", "", "hostname 默认为空")
	flag.IntVar(&port, "P", 0, "port 默认为 0 ")

	flag.Parse()
	fmt.Printf("user = %v \npwd = %v \nhost = %v \nport = %v",
		user, pwd, host, port)
}
