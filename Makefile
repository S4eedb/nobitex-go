WD=$(shell pwd)
init: setup
	mkdir -p build
setup:
	go mod download
	go mod tidy
generate:
	 go generate tools/api.gen.go
build_api: init
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o build/api $(WD)/cmd
	