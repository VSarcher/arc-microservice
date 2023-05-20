FROM golang:1.19

# RUN apk update && apk add --no-cache git

RUN mkdir /app

ADD . /app

WORKDIR /app

RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .

EXPOSE 8080

CMD ["/app/main"]