BINARY_NAME=main

build:
	go build -o $(BINARY_NAME)

run:
	go run $(BINARY_NAME).go

clean:
	rm -f $(BINARY_NAME)