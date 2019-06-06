package main

import (
	"io"
	"log"
	"net/http"
)

func handleHTTPProxy(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Proxy for request: [HTTP] %s - %s - %s", r.Method, r.Host, r.RequestURI)
	resp, err := http.DefaultTransport.RoundTrip(r)
	if err != nil {
		log.Printf("Cannot execute request. ERR: %v\n", err)
		http.Error(rw, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}
	defer resp.Body.Close()
	addSignatureHeader(rw.Header())
	copyHeader(rw.Header(), resp.Header)
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)
}
