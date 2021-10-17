.PHONY: build lint

build:
	@docker-compose build

logging: build
	@docker-compose up -d elasticsearch kibana

fluentbit: logging
	@docker-compose up -d fluentbit

run: fluentbit
	@docker-compose up -d sidelogger

down:
	@docker-compose down