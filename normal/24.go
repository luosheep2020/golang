package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(t)
	str := "This is an example of a string"
	fmt.Println(strings.HasPrefix(str,"This"))
	fmt.Println(strings.HasSuffix(str,"ing"))
	fmt.Println(strings.Contains(str,"an"))
	fmt.Println(strings.Replace(str,"an","a",-1))
}
