FROM golang:1.15.2-alpine

WORKDIR /go/src/calculator-service
COPY . .

RUN go build main.go

EXPOSE 9000
