FROM golang:1.19.0

WORKDIR /stock-service

RUN go install github.com/cosmtrek/air@latest

COPY . .
