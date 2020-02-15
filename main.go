package main

import (
	"fmt"
	"net"
)

func main() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")

	if err == nil {
		fmt.Println("Connection Successful!")
	}

	for i := 1; i <= 1024; i++ {
		address := fmt.Sprintf("scanme.nmap.org:%d", i)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			// port is closed or filtered.
			continue
		}
		conn.Close()
		fmt.Printf("%d open\n", i)
	}
}
