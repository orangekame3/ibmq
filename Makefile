.PHONY: install build

install:
	go install github.com/orangekame3/ibmq-cli@latest

build:
	go build -o ibmq

