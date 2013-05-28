GOPATH = $(PWD)
export GOPATH

.PHONY: bin/globol fmt

bin/globol:
	go build -o bin/globol globol

fmt:
	go fmt globol globol/lexer
