package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Open("./normal/output.txt")
	if err !=nil {
		fmt.Println("文件打开失败")
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	reader := bufio.NewReader(file)
	for  {
		str, err := reader.ReadString('\n')
		if err ==io.EOF {
			break
		}
		fmt.Print(str)
	}
	fmt.Println("文件读取结束...")
}
