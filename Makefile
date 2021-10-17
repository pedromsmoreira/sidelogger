.PHONY: build lint

build:
	@docker-compose build

logging: build
	@docker-compose up -d elasticsearch kibana fluentbit

filebeat: logging
	@docker-compose up -d filebeat

run:
	@docker-compose up -d sidelogger

down:
	@docker-compose down