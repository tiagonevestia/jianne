.PHONY: go_binary build

go_binary: 
	go build -o bin/jianne

build: go_binary