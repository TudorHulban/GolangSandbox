/*
	Echo Server.
*/
package main

import (
	"bufio"
	"io"
	"log"
	"net"
	"os"
)

const socket = "127.0.0.1:8081"

func handleConn(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		message, errReader := reader.ReadBytes(byte('\n'))
		if errReader != nil {
			if errReader != io.EOF {
				log.Println("failed to read data", errReader)
			}
			return
		}
		log.Println("request: ", string(message))

		conn.Write([]byte(message)) // writing back the message
	}
}

func main() {
	listener, errListen := net.Listen("tcp", socket)
	if errListen != nil {
		log.Println("listener error: ", errListen)
		os.Exit(1)
	}
	log.Println("listening on ", socket)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("failed to accept connection:", err)
			continue
		}
		go handleConn(conn)
	}
}
