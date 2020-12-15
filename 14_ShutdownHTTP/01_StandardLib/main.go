package main

// Example of HTTP Server Shutdown

import (
	"context"

	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

const socket = "127.0.0.1:8080"

func handleInterrupt(serverHTTP *http.Server, terminationSecs uint8) {
	sigint := make(chan os.Signal, 1)
	signal.Notify(sigint, os.Interrupt)

	<-sigint // the read from channel blocks until interrupt is received and sent on channel.

	// we can now shutdown
	log.Println("shutting down...")
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(int(terminationSecs))*time.Second)
	defer cancel()

	if errShutdown := serverHTTP.Shutdown(ctx); errShutdown != nil {
		log.Printf("HTTP server Shutdown: %v", errShutdown)
	}
}

func main() {
	httpServer := http.Server{
		Addr: socket,
	}

	go func() {
		log.Println("listening on ", socket)
		if errHTTP := httpServer.ListenAndServe(); errHTTP != http.ErrServerClosed {
			log.Fatalf("HTTP server ListenAndServe: %v", errHTTP)
		}
	}()
	handleInterrupt(&httpServer, 3)
}
