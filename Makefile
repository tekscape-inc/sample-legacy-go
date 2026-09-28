.PHONY: build clean

build:
	go build -o bin/legacy ./internal/app/legacy/cmd/legacy

clean:
	rm -rf bin
