build:
	mkdir -p bin
	go build -o bin/ip-controller cmd/ip-controller/main.go

run:
	go run cmd/ip-controller/main.go

test:
	go test ./...
