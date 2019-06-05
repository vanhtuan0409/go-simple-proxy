package main

import (
	"io/ioutil"
	"strings"

	"github.com/spf13/pflag"
)

type config struct {
	port           int
	blacklistHosts map[string]bool
}

func parseConfig() (*config, error) {
	var (
		fPort          = pflag.Int("port", 8080, "Proxy listen port")
		fBlacklistFile = pflag.String("blacklist-file", "", "List of forbidden hosts, one host per line")
	)

	pflag.Parse()

	blacklistHosts, err := readBlacklistToMap(*fBlacklistFile)
	if err != nil {
		return nil, err
	}

	return &config{
		port:           *fPort,
		blacklistHosts: blacklistHosts,
	}, nil
}

func readBlacklistToMap(p string) (map[string]bool, error) {
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")
	m := make(map[string]bool)

	for _, line := range lines {
		m[line] = true
	}

	return m, nil
}
