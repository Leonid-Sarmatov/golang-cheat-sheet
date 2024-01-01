package main

import (
	//"github.com/urfave/cli/v2"
	"fmt"
	//"os"
	//"log"
	"net"
	"time"
)

func Check(destinatior string, port string) string {
	address := destinatior + ":" + port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string
	if err != nil {
		status = fmt.Sprintf("[DOWN] %v is unreachable,\n Error: %v\n", destinatior, err)
	} else {
		status = fmt.Sprintf("[UP] %v is reachable,\n From: %v\n To: %v\n",
			destinatior, conn.LocalAddr(), conn.RemoteAddr())
	}
	return status
}
