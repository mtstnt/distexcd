.PHONY: build run-daemon run-cli clean

build:
	go build -o ./bin/ ./cmd/distexcd && \
	go build -o ./bin/ ./cmd/distexcd-cli

run-daemon:
	./bin/distexcd

run-cli:
	./bin/distexcd-cli

clean:
	rm -rf ./bin