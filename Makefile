deps:
	@glide up -v

run:
	@go run .

build:
	@go build -o bin/rtbutil

clean:
	@rm -Rf ./vendor ./bin

all: clean deps build