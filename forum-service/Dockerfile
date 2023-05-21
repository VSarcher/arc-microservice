FROM golang:1.19

WORKDIR usr/src/app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download
