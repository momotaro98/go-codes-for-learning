ROOT_DIR := $(dir $(realpath $(lastword $(MAKEFILE_LIST))))

build:
	cd lib/rustaceanize && cargo build --release
	cp lib/rustaceanize/target/release/librustaceanize.dylib lib/
	echo 'ROOT_DIR is $(ROOT_DIR)'
	go build -x -ldflags="-r $(ROOT_DIR)lib" main.go

run: build
	./main

