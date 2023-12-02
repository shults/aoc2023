test:
	@go test ./...

build:
	@mkdir -p bin
	@{ \
  		set -e ;\
  		FILES_TO_COMPILE=$$(ls *.go); \
  		for _file in $${FILES_TO_COMPILE}; do \
  		  CGO_ENABLED=0 go build -o bin/$$(basename $${_file} .go) $${_file}; \
	  	done \
  	}