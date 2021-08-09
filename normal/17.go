package main

import (
	"fmt"

)

func main() {
	list:=[...]int{1,2,3}
	for i,v:=range list{
		fmt.Printf("%d %d\n",i,v )
	}
}
