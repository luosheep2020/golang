package main

import (
	"fmt"
	"time"
)

func says(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100*time.Microsecond)
		fmt.Println(s)
	}
}

func main() {
	go says("Hello")
	says("World")
}
