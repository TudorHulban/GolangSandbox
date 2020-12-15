package main

import (
	"log"
	"net"
	"os"
)

const socket = "127.0.0.1:8081"
const ramDisk = ""
const cmd = "firefox -headless --screenshot "

func init() {
	runCmd("bash", "cleanpng.sh")
}

func main() {
	listener, errListen := net.Listen("tcp", socket)
	if errListen != nil {
		log.Println("listener error: ", errListen)
		os.Exit(1)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("failed to accept conn:", err)
			continue
		}
		handleConn(conn) // could be also concurrent.
	}
}
