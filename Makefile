.PHONY: build/pager-client

build/pager-client:
	mkdir -p ./build
	go build -o ./build/pager-client main.go

build: build/pager-client

install:
	mkdir -p ~/.local/bin
	cp ./build/pager-client ~/.local/bin/pager-client
