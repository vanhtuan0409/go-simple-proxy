package main

import (
	"io"
	"net/http"
)

func copyHeader(dst, src http.Header) {
	for k, vv := range src {
		for _, v := range vv {
			dst.Add(k, v)
		}
	}
}

func pipe(dst io.WriteCloser, src io.ReadCloser) {
	defer dst.Close()
	defer src.Close()
	io.Copy(dst, src)
}

func pipeWithTee(dst io.WriteCloser, src io.ReadCloser, sniffer io.Writer) {
	defer dst.Close()
	defer src.Close()

	r := io.TeeReader(src, sniffer)
	io.Copy(dst, r)
}

func addSignatureHeader(h http.Header) {
	h.Set("X-Proxy-Name", "tuanvuong")
}
