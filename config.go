package main

import (
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/spf13/pflag"
)

type config struct {
	port           int
	blacklistHosts *regexp.Regexp
}

func parseConfig() (*config, error) {
	var (
		fPort          = pflag.Int("port", 8080, "Proxy listen port")
		fBlacklistFile = pflag.String("blacklist-file", "", "List of forbidden hosts, one host per line")
	)

	pflag.Parse()

	pattern, err := readBlacklistToPattern(*fBlacklistFile)
	if err != nil {
		return nil, err
	}

	return &config{
		port:           *fPort,
		blacklistHosts: pattern,
	}, nil
}

func readBlacklistToPattern(p string) (*regexp.Regexp, error) {
	content, err := ioutil.ReadFile(p)
	if err != nil {
		return nil, err
	}

	normalize := strings.TrimSpace(string(content))
	lines := strings.Split(normalize, "\n")
	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}
	pattern := strings.Join(lines, "|")
	rx, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}

	return rx, nil
}
