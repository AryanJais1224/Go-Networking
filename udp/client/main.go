package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	addr, err := net.ResolveUDPAddr("udp", "localhost:8082")
	if err != nil {
		panic(err)
	}

	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')

		conn.Write([]byte(text))

		buffer := make([]byte, 1024)
		n, _, _ := conn.ReadFromUDP(buffer)
		fmt.Println(string(buffer[:n]))
	}
}
