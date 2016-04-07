all: deps build

deps:
	go get -u github.com/spf13/cobra

build:
	go build

