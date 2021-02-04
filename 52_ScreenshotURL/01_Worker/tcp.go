package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func handleConn(conn *net.Conn) {
	defer (*conn).Close()

	reader := bufio.NewReader(*conn)
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

		resp, errWork := doWork(request)
		if errWork != nil {
			if _, errWrite := (*conn).Write([]byte("try later" + "\n")); errWrite != nil {
				log.Println("could not write on error:", errWrite)
			}
			continue
		}

		if _, errWrite := (*conn).Write([]byte(resp + "\n")); errWrite != nil {
			log.Println("could not write response:", errWrite)
		}
		break // faster of course without closing and opening. chose to close conn for mock up
	}
}
