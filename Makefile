all: tools deps build lint

tools:
	go get -u github.com/golang/lint/golint
	go get -u github.com/Masterminds/glide

deps:
	glide install

build:
	go build

lint:
	golint

clean:
	rm -f ./rave
