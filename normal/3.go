package main

import "fmt"

func Fac(n uint64) (result uint64) {
	if n>0 {
		result =n*Fac(n-1)
		return result
	}
	return  1
}
func main() {
	num:=15
	fmt.Printf("%d的阶乘是 %d \n",num,Fac(uint64(num)))
}
