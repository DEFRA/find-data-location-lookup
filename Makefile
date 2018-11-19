

all: test run

test:
	@go test

run:
	@go run *.go

build:
	@go build -o ll.exe *.go