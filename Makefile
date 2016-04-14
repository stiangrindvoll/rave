all: tools deps build lint

tools:
	go get -u github.com/golang/lint/golint

deps:
	go get -u github.com/spf13/cobra

build:
	go build

lint:
	golint

clean:
	rm -f ./rave
