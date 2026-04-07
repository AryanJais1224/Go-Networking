package main

import (
	"io"
	"net"
	"sync/atomic"
)

var backends = []string{
	"localhost:9001",
	"localhost:9002",
}

var counter uint32

func nextBackend() string {
	i := atomic.AddUint32(&counter, 1)
	return backends[int(i)%len(backends)]
}

func handle(conn net.Conn) {
	defer conn.Close()
	backendConn, err := net.Dial("tcp", nextBackend())
	if err != nil {
		return
	}
	defer backendConn.Close()

	go io.Copy(backendConn, conn)
	io.Copy(conn, backendConn)
}

func main() {
	ln, _ := net.Listen("tcp", ":8080")
	for {
		conn, _ := ln.Accept()
		go handle(conn)
	}
}
