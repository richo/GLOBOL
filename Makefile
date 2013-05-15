GOPATH = $(PWD)
export GOPATH

.PHONY: bin/globol

bin/globol:
	go build -o bin/globol globol
