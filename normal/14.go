package main

import (
	"fmt"
)

func getDate()(int ,int)  {
	return 100,200
}
func main() {
	a,_ :=getDate()
	_,b :=getDate()
	fmt.Print(a,b,"\n")

}
