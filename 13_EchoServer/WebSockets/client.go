package main

import (
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

const urlBinance = "wss://stream.binance.com:9443/ws/bnb@kline_1m"

func webSocketConnection(requestHeader http.Header, url *url.URL) (*websocket.Conn, *http.Response, error) {
	log.Printf("connecting to %s", url.String())

	return websocket.DefaultDialer.Dial(url.String(), nil)
}

func readingMessages(connection *websocket.Conn, stop chan struct{}) {
	defer close(stop)

	for {
		_, message, err := connection.ReadMessage()
		if err != nil {
			log.Println("read glitch:", err)
			return
		}

		go log.Printf("received message: %s", message)
	}
}

func sendingMessages(connection *websocket.Conn, messages <-chan string, interrupt chan os.Signal, stop chan struct{}) {
	for {
		select {
		case <-stop:
			{
				return
			}

		case msg := <-messages:
			{
				err := connection.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					log.Println("write:", err)
					return
				}
			}

		case <-interrupt:
			{

				log.Println("interrupt")

				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				errClosing := connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if errClosing != nil {
					log.Println("closing communication:", errClosing)
					return
				}

				stop <- struct{}{}

				return
			}
		}
	}
}

func main() {
	// u := url.URL{
	// 	Scheme: "ws",
	// 	Host:   "localhost:7000",
	// 	Path:   "/echo",
	// }

	u, errParse := url.Parse(urlBinance)
	if errParse != nil {
		log.Println("read:", errParse)
		os.Exit(1)
	}

	conn, _, errConnect := webSocketConnection(nil, u)
	if errConnect != nil {
		log.Println("read:", errConnect)
		os.Exit(2)
	}

	data := make(chan string)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	stop := make(chan struct{})

	go readingMessages(conn, stop)

	ticker := time.NewTicker(time.Second * 3)
	defer ticker.Stop()

	go sendingMessages(conn, data, interrupt, stop)

	data <- "started"

	<-stop
}
