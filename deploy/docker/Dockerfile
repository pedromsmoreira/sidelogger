FROM golang:1.16 as build

WORKDIR /app
COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN go test -v -timeout 30s github.com/pedromsmoreira/sidelogger/loggers
RUN go build -o /sidelogger

##
## Deploy
##
FROM ubuntu:focal-20210921 as runtime
WORKDIR /
COPY --from=build /sidelogger /sidelogger
COPY --from=build /app/config.yaml config.yaml
EXPOSE 5000
ENTRYPOINT [ "/sidelogger" ]