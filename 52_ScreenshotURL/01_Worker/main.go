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
	errInit := runCmd("bash", "cleanpng.sh")
	if errInit != nil && errInit.Error() != "exit status 1" {
		log.Println("clean images error:", errInit)
	}
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
		handleConn(&conn) // no goroutines due to firefox limitations
		log.Println("---------- request handled")
	}
}
