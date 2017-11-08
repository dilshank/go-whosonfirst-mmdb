CWD=$(shell pwd)
GOPATH := $(CWD)

prep:
	if test -d pkg; then rm -rf pkg; fi

self:   prep rmdeps
	if -d src; then rm -rf src; fi
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-mmdb/
	cp *.go src/github.com/whosonfirst/go-whosonfirst-mmdb/
	mkdir -p src/github.com/whosonfirst/go-whosonfirst-mmdb/provider
	cp provider/*.go src/github.com/whosonfirst/go-whosonfirst-mmdb/provider/
	cp -r vendor/* src/

rmdeps:
	if test -d src; then rm -rf src; fi 

build:	fmt bin

# if you're wondering about the 'rm -rf' stuff below it's because Go is
# weird... https://vanduuren.xyz/2017/golang-vendoring-interface-confusion/
# (20170912/thisisaaronland)

deps:
	@GOPATH=$(GOPATH) go get -u "github.com/oschwald/maxminddb-golang"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-csv"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-geojson-v2"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-iplookup"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-log"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-spr"
	@GOPATH=$(GOPATH) go get -u "github.com/whosonfirst/go-whosonfirst-uri"
	rm -rf src/github.com/whosonfirst/go-whosonfirst-iplookup/vendor/github.com/whosonfirst/go-whosonfirst-spr
	rm -rf src/github.com/whosonfirst/go-whosonfirst-geojson-v2/vendor/github.com/whosonfirst/go-whosonfirst-flags

vendor-deps: rmdeps deps
	if test -d vendor; then rm -rf vendor; fi
	mkdir vendor
	cp -r src/* vendor/
	find vendor -name '.git' -print -type d -exec rm -rf {} +
	rm -rf src
	rm -rf vendor/github.com/oschwald/maxminddb-golang/test-data

fmt:
	go fmt *.go
	go fmt cmd/*.go

bin: 	self
	@GOPATH=$(GOPATH) go build -o bin/wof-mmdb cmd/wof-mmdb.go
	@GOPATH=$(GOPATH) go build -o bin/wof-mmdb-prepare cmd/wof-mmdb-prepare.go
	@GOPATH=$(GOPATH) go build -o bin/wof-mmdb-server cmd/wof-mmdb-server.go
