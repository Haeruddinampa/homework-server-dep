install:
	go mod tidy

lint:
	go fmt ./...

test: install
	go test -v

build: install
	go build

run: install
	./homework-server-depl.exe