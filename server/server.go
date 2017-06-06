package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	service := "127.0.0.1:11066"

	listener, err := net.Listen("tcp", service)
	checkError(err)

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		go handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("client")
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			return
		}
		fmt.Println("read ", string(buf[0:n]))
		_, err2 := conn.Write(buf[0:n])
		if err2 != nil {
			return
		}
		fmt.Println("write")
	}
}

func checkError(err error) {
	if err != nil {
		fmt.Println("fatalerror", err.Error())
		os.Exit(1)
	}
}
