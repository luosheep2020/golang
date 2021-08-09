package main

import (
	"fmt"
	"runtime"
	"sync"
)

var  count int = 0

func Count(locker *sync.Mutex)  {
	locker.Lock()
	count++
	fmt.Println(count)
	locker.Unlock()
}
func main() {
	lock := &sync.Mutex{}
	for i := 0; i < 10; i++ {
		go  Count(lock)
	}
	for  {
		lock.Lock()
		c := count
		lock.Unlock()
		runtime.Gosched()
		if  c>=10{
			break
		}
	}
}
