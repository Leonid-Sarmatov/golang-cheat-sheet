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
	adress := destinatior+":"+port
	timeout := time.Duration(5 * time.Second)
	conn, err := net.DialTimeout("tcp", address, timeout)

	var status string
	if err != nil {
		status = fmt.Sprint
	}
}