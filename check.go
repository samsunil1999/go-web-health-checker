package main

import (
	"fmt"
	"net"
	"time"
)

func Check(destination string, port string) string {
	address := destination + ":" + port
	timeout := time.Duration(5 * time.Second)

	var status string

	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable, \n Error: %v", destination, err)
	} else {
		status = fmt.Sprintf("[UP] %v is running, \n From: %v\n To: %v\n",
			destination,
			conn.LocalAddr(),
			conn.RemoteAddr())
	}

	return status
}
