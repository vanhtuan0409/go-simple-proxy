package main

import (
	"log"
	"net"
	"net/http"
)

func handleConnectProxy(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Proxy for request: [Connect] %s - %s - %s", r.Method, r.Host, r.RequestURI)
	dialer := net.Dialer{}
	outConn, err := dialer.Dial("tcp", r.Host)
	if err != nil {
		log.Printf("Cannot dial TCP connection. ERR: %v\n", err)
		http.Error(rw, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	rw.WriteHeader(http.StatusOK)
	addSignatureHeader(rw.Header())
	hijacker, ok := rw.(http.Hijacker)
	if !ok {
		log.Printf("Hijacking not supported. ERR: %v", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	inConn, _, err := hijacker.Hijack()
	if err != nil {
		log.Printf("Cannot hijack TCP connection. ERR: %v\n", err)
		http.Error(rw, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	go pipe(outConn, inConn)
	go pipe(inConn, outConn)
}
