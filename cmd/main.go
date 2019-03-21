package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
)

type defaultconfig struct {
	address string
	port    string
}

var dc = defaultconfig{
	address: "127.0.0.1",
	port:    "8000",
}

func main() {
	fmt.Println("Start a proxy instance ...")
	go tcpProxy(dc)
	admHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "admin page\n")
	}
	http.HandleFunc("/", admHandler)
	http.ListenAndServe(":9001", nil)
}

func tcpProxy(dc defaultconfig) {
	lis, derr := net.Listen("tcp", dc.address+":"+dc.port)
	if derr != nil {
		fmt.Println("listener error ", derr)
	}
	defer lis.Close()

	conn, err := lis.Accept()
	if err != nil {
		fmt.Println("error of accept", err)
	}

	for {

		rconn, err := net.Dial("tcp4", "127.0.0.1:9002")
		if err != nil {
			fmt.Println("error of dial", err)
		}
		defer rconn.Close()

		var buf = make([]byte, 10)
		var buff = make([]byte, 10)
		read, err := conn.Read(buf)
		if err != nil {
			fmt.Println("error of read ccon", err.Error)
			break
		}
		b := buf[:read]
		rconn.Write(b)

		write, _ := rconn.Read(buff)
		c := buff[:write]
		conn.Write(c)
	}
}
