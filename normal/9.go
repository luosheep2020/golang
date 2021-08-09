package main

import "fmt"

func main() {
	var map1 map[string]string
	map1= make(map[string]string)

	map1["France"]="巴黎"
	map1["Japan"]="东京"
	map1["Canada"]="多伦多"
	map1["Italy"]="罗马"

	for country := range map1{
		fmt.Println(country,"首都是",map1[country])
	}

	capital, ok := map1["America"]
	if ok {
		fmt.Println("America的首都是",capital	)
	}else {
		fmt.Println("不存在该元素")
	}
}