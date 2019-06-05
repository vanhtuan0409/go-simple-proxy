package main

import "github.com/spf13/pflag"

type config struct {
	port int
}

func parseConfig() config {
	var (
		fPort = pflag.Int("port", 8080, "Proxy listen port")
	)

	pflag.Parse()

	return config{
		port: *fPort,
	}
}
