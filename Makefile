.PHONY: build lint

build:
	@docker-compose build

logging: build
	@docker-compose up -d logstash elasticsearch kibana 

fluentbit: logging
	@docker-compose up -d fluentbit

run: 
	@docker-compose up -d sidelogger

down:
	@docker-compose down