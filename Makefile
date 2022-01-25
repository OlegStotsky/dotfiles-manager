.PHONY: build resources

all: generate build

build:
	goimports -w .
	go mod tidy
	go install .

generate: resources

resources:
	go-bindata -o pkg/resources/bindata.go -prefix resources -pkg resources resources/...
