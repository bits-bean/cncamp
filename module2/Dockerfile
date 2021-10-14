# syntax=docker/dockerfile:1

FROM golang:1.17-alpine AS builder
WORKDIR /app

COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /server

# 采用多阶段构建，使原来镜像的大小从321MB缩小到11MB
FROM alpine:latest

COPY --from=builder /server /

EXPOSE 8000

CMD [ "/server" ]