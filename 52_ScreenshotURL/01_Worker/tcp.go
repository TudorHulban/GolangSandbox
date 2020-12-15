package main

import (
	"bufio"
	"io"
	"log"
	"net"
)

func handleConn(pConnection *net.Conn) {
	defer (*pConnection).Close()

	reader := bufio.NewReader(*pConnection)
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
			_, errWrite := (*pConnection).Write([]byte("try later" + "\n"))
			if errWrite != nil {
				log.Println("could not write on error:", errWrite)
			}
			continue
		}
		_, errWrite := (*pConnection).Write([]byte(resp + "\n"))
		if errWrite != nil {
			log.Println("could not write response:", errWrite)
		}
		break // faster of course without closing and opening. chose to close conn for mock up
	}
}
