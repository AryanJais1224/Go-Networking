package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("Enter message: ")
		text, _ := reader.ReadString('\n')
		conn.Write([]byte(text))

		reply := make([]byte, 1024)
		n, _ := conn.Read(reply)
		fmt.Println(string(reply[:n]))
	}
}
