.PHONY: build

build: up
	go build -o go-game .

up:
	docker-compose up -d