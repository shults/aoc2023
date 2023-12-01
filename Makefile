test:
	@go test ./...

build:
	@mkdir -p build
	@go build -o build/t01 t01.go