build:
	go build -o bin/pikachu .

run:
	go run *.go --blacklist-file blacklists.txt

start-browser:
	/opt/google/chrome/chrome --proxy-server=http://localhost:8080

sample-http:
	curl -v -x http://localhost:8080 google.com


sample-tunnel:
	curl -vpx http://localhost:8080 google.com
