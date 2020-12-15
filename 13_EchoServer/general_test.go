package main

import (
	"bufio"
	"fmt"
	"net"
	"testing"

	"github.com/stretchr/testify/assert"
)

const payload = "abcd"

func TestListen(t *testing.T) {
	conn, errDial := net.Dial("tcp", socket)

	if assert.NoError(t, errDial, "could not connect") {
		message := payload + "\n"

		fmt.Fprintf(conn, message) // sends message through the provided io.writer

		reply, errRead := bufio.NewReader(conn).ReadString('\n')
		assert.NoError(t, errRead)
		assert.Equal(t, message, reply)
	}

}
