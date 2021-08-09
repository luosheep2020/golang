package main

import "fmt"

func main() {
	var a int =20
	var ip *int
	ip = &a
	fmt.Printf("a的地址为:%x\n\n", &a)
	fmt.Printf("ip变量存放的地址为:%x\n\n", ip)
	fmt.Printf("ip变量的值为:%d\n\n",*ip)
}