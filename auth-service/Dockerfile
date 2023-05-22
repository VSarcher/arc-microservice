FROM golang:1.19

# RUN apk update && apk add --no-cache git

# RUN mkdir /app/auth

WORKDIR app

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go mod download

EXPOSE 3000

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


# CMD ["/app/main"]