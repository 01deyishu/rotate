package main

import (
	"fmt"
	"io"
	"net"
	"net/http"
	"time"
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
	go Listen(dc)
	admHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "admin page\n")
	}
	http.HandleFunc("/admin", admHandler)
	http.ListenAndServe(":9001", nil)
}

func Listen(dc defaultconfig) {
	lis, err := net.Listen("tcp", dc.address+":"+dc.port)
	if err != nil {
		fmt.Println("listener error ", err)
	}
	go handleconn(lis)
	time.Sleep(time.Second * 20)
}

func handleconn(l net.Listener) {
	fmt.Println(l.Addr().Network())
}
