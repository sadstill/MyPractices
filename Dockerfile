FROM golang:1.23-alpine AS builder

COPY . /app
WORKDIR /app

RUN go mod download
RUN go build -o ./bin/chat-server cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /app/bin/chat-server .

CMD ["./chat-server"]
