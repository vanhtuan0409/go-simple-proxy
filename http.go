package main

import (
	"io"
	"log"
	"net/http"
)

func handleHTTPProxy(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Proxy for request: [%s] %s - %s", r.URL.Scheme, r.Method, r.RequestURI)
	client := http.Client{}
	req, err := http.NewRequest(r.Method, r.RequestURI, r.Body)
	if err != nil {
		log.Printf("Cannot create request. ERR: %v\n", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	copyHeader(req.Header, r.Header)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("Cannot execute request. ERR: %v\n", err)
		http.Error(rw, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()
	rw.Header().Set("X-Proxy-Name", "tuanvuong")
	copyHeader(rw.Header(), resp.Header)
	rw.WriteHeader(resp.StatusCode)
	io.Copy(rw, resp.Body)
}
