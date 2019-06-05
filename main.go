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

type proxy struct {
	server         *http.Server
	blacklistHosts map[string]bool
}

func main() {
	config, err := parseConfig()
	if err != nil {
		log.Fatalf("Cannot parse config file. ERR: %v", err)
	}

	p := createProxy(config)
	errChan := make(chan error, 2)
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt)

	go func() {
		errChan <- p.server.ListenAndServe()
	}()

	go func() {
		sig := <-sigChan
		log.Printf("Received signal: %v. Shutting down server\n", sig)
		ctx, cancelFn := context.WithTimeout(context.Background(), 15*time.Second)
		defer cancelFn()
		errChan <- p.server.Shutdown(ctx)
	}()

	log.Printf("Server is listenning on port :%d", config.port)

	err = <-errChan
	if err != nil && err != http.ErrServerClosed {
		log.Fatalf("Cannot start or shutdown server. ERR: %v\n", err)
	}

	log.Println("Server shutdown. Glad to work with you")
}

func createProxy(conf *config) *proxy {
	addr := fmt.Sprintf(":%d", conf.port)
	p := &proxy{}
	p.blacklistHosts = conf.blacklistHosts
	server := &http.Server{
		Addr:         addr,
		Handler:      http.HandlerFunc(p.handler),
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
	}
	p.server = server
	return p
}

func (p *proxy) handler(rw http.ResponseWriter, r *http.Request) {
	if p.blacklistHosts[r.Host] {
		http.Error(rw, "Forbidden", http.StatusForbidden)
		return
	}

	if r.Method == http.MethodConnect {
		handleConnectProxy(rw, r)
	} else {
		handleHTTPProxy(rw, r)
	}
}
