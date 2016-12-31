default: build

bin/redisutil: *.go
	go build -race -v -o $@ $^

build: bin/redisutil

deps:
	go get github.com/garyburd/redigo/redis
