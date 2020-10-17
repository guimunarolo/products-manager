FROM golang:1.15.2-alpine

WORKDIR /go/src/calculator-service
COPY . .

RUN apk add build-base
RUN go build main.go

EXPOSE 9000
