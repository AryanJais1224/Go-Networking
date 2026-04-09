package main

import (
	"fmt"
	"net"
	"sync"
	"time"
)

func scan(port int, wg *sync.WaitGroup) {
	defer wg.Done()
	address := fmt.Sprintf("localhost:%d", port)
	conn, err := net.DialTimeout("tcp", address, time.Millisecond*200)
	if err == nil {
		fmt.Println("open", port)
		conn.Close()
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 1024; i++ {
		wg.Add(1)
		go scan(i, &wg)
	}
	wg.Wait()
}
