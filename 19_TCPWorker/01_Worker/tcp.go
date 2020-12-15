package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func handleConn(connection net.Conn) {
	defer connection.Close()

	reader := bufio.NewReader(connection)
	for {
		bytes, errReader := reader.ReadBytes(byte('\n'))
		if errReader != nil {
			if errReader != io.EOF {
				log.Println("failed to read data", errReader)
			}
			return
		}

		request := string(bytes)
		log.Println("request: ", request)

		resp, err := doWork(request)
		if err != nil {
			connection.Write([]byte("try later" + "\n"))
			continue
		}
		connection.Write([]byte(resp + "\n"))
	}
}
