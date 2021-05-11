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

func socksConnection(requestHeader http.Header) (*websocket.Conn, *http.Response, error) {
	u := url.URL{
		Scheme: "ws",
		Host:   "localhost:8080",
		Path:   "/echo",
	}

	log.Printf("connecting to %s", u.String())

	return websocket.DefaultDialer.Dial(u.String(), nil)
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

func sendingMessages(connection *websocket.Conn, data chan string, interrupt chan os.Signal, stop chan struct{}) {
	for {
		select {
		case <-stop:
			{
				return
			}

		case msg := <-data:
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
	conn, _, errConnect := socksConnection(nil)
	if errConnect != nil {
		log.Println("read:", errConnect)
		os.Exit(1)
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
