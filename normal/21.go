package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
  filePath:= "./output.txt"
  file,err:=os.OpenFile(filePath,os.O_WRONLY|os.O_CREATE,0666)
	if err!=nil {
		fmt.Printf("打开文件错误=%v\n",err)
		return
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	str:="golang开发\r\n"
	writer := bufio.NewWriter(file)
	for i := 0; i < 3; i++ {
		_, err := writer.WriteString(str)
		if err != nil {
			return
		}
	}
	err = writer.Flush()
	if err != nil {
		return
	}
}
