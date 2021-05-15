package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"

	"github.com/gorilla/websocket"
)

type Config struct {
	RequestHeader       http.Header
	URI                 string
	PongIntervalSeconds uint
}

// Client Concentrates websocket information.
type Client struct {
	connection *websocket.Conn
	URL        url.URL

	send      chan []byte
	stop      chan struct{}
	interrupt chan os.Signal
}

const urlBinance = "wss://stream.binance.com:9443/ws/bnbusdt@trade"

func NewClient(cfg Config) (*Client, error) {
	url, errParse := url.Parse(cfg.URI)
	if errParse != nil {
		return nil, errParse
	}

	conn, _, errConn := websocket.DefaultDialer.Dial(url.String(), cfg.RequestHeader)
	if errConn != nil {
		return nil, errConn
	}

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	return &Client{
		connection: conn,
		URL:        *url,
		send:       make(chan []byte),
		stop:       make(chan struct{}),
		interrupt:  interrupt,
	}, nil
}

func (c *Client) ReadMessages() {
	defer c.cleanUp()

	var shouldStop bool

	checkSignals := func() {
		select {
		case <-c.interrupt:
			{
				log.Println("interrupt")
				shouldStop = true
			}
		}
	}

	go checkSignals()

	for !shouldStop {
		_, message, errRead := c.connection.ReadMessage()
		if errRead != nil {
			log.Println("read glitch:", errRead)
			return
		}

		go log.Printf("received message: %s", message)

	}

	c.stop <- struct{}{}
}

func (c *Client) cleanUp() {
	close(c.stop)
	close(c.send)
}

func main() {
	cfg := Config{
		URI: urlBinance,
	}

	c, errNew := NewClient(cfg)
	if errNew != nil {
		log.Println(errNew)
		os.Exit(1)
	}

	go c.ReadMessages()

	<-c.stop
}
