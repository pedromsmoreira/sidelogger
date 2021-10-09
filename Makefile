.PHONY: build lint

build:
	@docker-compose build

logging: build
	@docker-compose up -d elasticsearch kibana logstash

filebeat: logging
	@docker-compose up -d filebeat

run: filebeat
	@docker-compose up -d sidelogger

down:
	@docker-compose down