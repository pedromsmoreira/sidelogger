# Sidelogger

A simple PoC to test multiple logging libraries in Go it also implements all the needed logging infrastructure in a containerized way.
This PoC logs all the lines to `docker stdout` which is being watched by a filebeat container. 
The filebeat will be scraping the `/var/lib/docker/containers` files in order to fetch the entries and send them to a Logstash pipeline which will filter and insert the entries in an Elastic Search index. This index will be queried by Kibana.

# Tech specifications

Language used: Go
Logging Stack: ELK

Future Improvements: Switch Filebeat and Logstash with FluentBit

# Use a different logger implementation
To change the logger implementation it is done with a simple change in the `config.yaml` file.
```yaml
logger:
  name: "logrus" # uberzap, logrus, apex
```


## What you need to run the PoC

- Docker
- Docker Compose
- Make

Command to execute:
```shell
make run
```

Note: might need to be ran twice if elasticsearch container takes to long to start up and makes Logstash container to terminate.


If you don't want to use Docker, you will need to install Go 1.16 or above.

Command to execute:
```shell
go run main.go
```
