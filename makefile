run:
	go build -o bin/main main.go && ./bin/main

test:
	go test ./...