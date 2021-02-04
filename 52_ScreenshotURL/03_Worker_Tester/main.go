package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const socket = "127.0.0.1:8081"

func main() {
	conn, _ := net.Dial("tcp", socket)

	toSend := "https://www.olx.ro/oferta/lifepo4-acumulator-baterie-incarcator-60ah-100ah-200ah-ID9XZiR.html"
	fmt.Fprintf(conn, toSend+"\n")

	reply, _ := bufio.NewReader(conn).ReadString('\n')
	log.Println("received: ", reply)
}
