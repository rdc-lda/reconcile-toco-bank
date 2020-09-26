deps:
	@glide up -v

run:
	@go run .

build:
	@go build -o bin/rtbutil

clean:
	@rm -Rf ./vendor ./bin

install: build
	@cp bin/* ~/Bin

all: clean deps build install