FROM golang:1.19

# RUN mkdir app/forum

WORKDIR app

RUN go install github.com/cosmtrek/air@latest

COPY . .
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux go build -o main .
# CMD ["/app/main"]
