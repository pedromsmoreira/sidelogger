.PHONY: build lint

build:
	@docker-compose build

run: build
	@docker-compose up -d sidelogger