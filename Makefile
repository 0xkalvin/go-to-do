.PHONY: dev build run all clean database down test

default: dev

build:
	go build -o output.out ./src/server

run:
	./output.out

all:
	@docker-compose up

dev:
	go run src/server/main.go

clean:
	rm -rf *.out

database:
	@docker-compose up -d postgres

down:
	@docker-compose down 

test:
	go test ./src/server -v