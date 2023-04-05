package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"
)

func main() {
	targetIP := ""
	var openPorts []int

	var wg sync.WaitGroup

	for port := 1; port <= 65535; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := targetIP + ":" + strconv.Itoa(port)
			conn, err := net.DialTimeout("tcp", address, 1*time.Second)
			if err == nil {
				openPorts = append(openPorts, port)
				conn.Close()
			}
		}(port)
	}

	wg.Wait()
	fmt.Printf("Open ports on %s: %v\n", targetIP, openPorts)
}
