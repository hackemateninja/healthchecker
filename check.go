package main

import (
	"fmt"
	"net"
	"time"
)

func Check(domain, port string) string {

	address := domain + ":" + port

	conn, err := net.DialTimeout("tcp", address, time.Duration(5*time.Second))

	if err != nil {
		return fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", domain, err)
	}

	return fmt.Sprintf("[UP] %v is reachable, \n from: %v\n To: %v", domain, conn.LocalAddr(), conn.RemoteAddr())
}
