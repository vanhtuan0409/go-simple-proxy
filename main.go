package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	port := 8080
	server := getServer(port)
	errChan := make(chan error, 2)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		errChan <- server.ListenAndServe()
	}()

	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %v. Shutting down server\n", sig)
		ctx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFn()
		errChan <- server.Shutdown(ctx)
	}()

	log.Printf("Server is listenning on port :%d", port)

	err := <-errChan
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Cannot start or shutdown server. ERR: %v\n", err)
	}

	log.Println("Server shutdown. Glad to work with you")
}

func getServer(port int) *http.Server {
	addr := fmt.Sprintf(":%d", port)
	return &http.Server{
		Addr:         addr,
		Handler:      http.HandlerFunc(handleHTTPProxy),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
}
