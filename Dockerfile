# 使用官方 Golang 镜像作为构建阶段
FROM golang:1.21.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o myapp .

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/myapp .
COPY --from=builder /app/config.yml .

CMD ["./myapp"]
