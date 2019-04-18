FROM golang:1.12-alpine as build-base
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh make 
ADD . /go/src/github.com/sah4ez/go-kafka-example
WORKDIR /go/src/github.com/sah4ez/go-kafka-example
RUN GOOS=linux GOARCH=amd64 GO111MODULE=on go build -o /tmp/app .

FROM alpine:latest
COPY --from=build-base /tmp/app /usr/bin/app
EXPOSE 9092 9093 9094 9095
VOLUME "/tmp/"
