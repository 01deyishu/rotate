package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	fmt.Println("Start a echo server to test ... ")
	l, err := net.Listen("tcp", "127.0.0.1:9002")
	if err != nil {
		fmt.Println("error of listen port ", err)
		os.Exit(1)
	}
	defer l.Close()
	fmt.Println("Listen port 9002 ... ")

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("error of accept data", err)
			os.Exit(1)
		}
		fmt.Printf("Receiving data from %s -> %s \n", conn.RemoteAddr(), conn.LocalAddr())

		go func(conn net.Conn) {
			defer conn.Close()
			for {
				io.Copy(conn, conn)
			}
		}(conn)
	}
}
