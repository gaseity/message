package main

import (
	"fmt"
	//"io"
	"net"
	"os"
)

func main() {
	//	server := os.Args[1]
	server := "127.0.0.1:11066"

	tcpAddr, err := net.ResolveTCPAddr("tcp4", server)
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	checkError(err)

	conn.Write([]byte("abab"))
	readFully(conn)
}

func checkError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Fatal error: %s", err.Error())
		os.Exit(1)
	}
}

func readFully(conn net.Conn) {
	defer conn.Close()

	var buf [512]byte
	for {
		n, _ := conn.Read(buf[0:])
		fmt.Println(string(buf[0:n]))
	}
}
