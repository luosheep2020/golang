package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- i
			time.Sleep(time.Second)
		}
	}()

	for data := range ch {
		fmt.Println(data)
		if data ==2{
			break
		}
	}
}
