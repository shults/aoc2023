test:
	@go test ./...

build:
	@mkdir -p bin
	@CGO_ENABLED=0 go build -o bin/aoc2023 main.go