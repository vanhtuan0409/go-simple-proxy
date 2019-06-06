# Go Simple Proxy

Simple proxy that supports:
- [x] HTTP proxy
- [x] HTTPS proxy via CONNECT
- [x] Backlist by domains

### Installation

```sh
go get -u github.com/vanhtuan0409/go-simple-proxy
```

Or if you want to build from source

```sh
git clone github.com/vanhtuan0409/go-simple-proxy go-simple-proxy
cd go-simple-proxy
make build
```

### Usage

```sh
pikachu --port 8080 --blacklist-file blacklists.txt
```
