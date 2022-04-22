.PHONY: build run clean all

build:
	go build -o bin/main util.go main.go

run:
	go run src/*.go

clean:
	rm -rf savedFiles
	rm -rf downloadFiles

all: build run


