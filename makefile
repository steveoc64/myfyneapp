all: wasm

run: build
	./myfyneapp

build:
	go build .

wasm:
	GOOS=js GOARCH=wasm go build .
	ls -ltra
