.PHONY: dev build run all clean database down

default: dev

build:
	go build -o output.out ./src

run:
	./output.out

all:
	@docker-compose up

dev:
	go run src/main.go

clean:
	rm -rf *.out

database:
	@docker-compose up -d postgres

down:
	@docker-compose down 