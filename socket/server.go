package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	checkError1(err)
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError1(err)
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		b := make([]byte, 1024)
		conn.Read(b)
		conn.Write([]byte(string(b) + ":" + strconv.FormatInt(time.Now().Unix(), 10)))
		conn.Close()
	}
}

func checkError1(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}