package main

import "fmt"

func main()  {
	var a =200
	var b  =100
	var result=max(a,b)
	fmt.Printf("最大值是：%d", result)
}

func max(a int, b int) int {
	var result int
	if a>b {
		result=a
	}else {
		result=b}
	return result
}
