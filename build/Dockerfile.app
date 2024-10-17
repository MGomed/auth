FROM golang:1.22-alpine AS builder

COPY . /github.com/MGomed/auth
WORKDIR /github.com/MGomed/auth

RUN go mod download
RUN go build -o ./bin/app cmd/main.go

FROM alpine:latest

WORKDIR /root/

RUN mkdir -p ./out/log

COPY --from=builder /github.com/MGomed/auth/bin/app .
COPY --from=builder /github.com/MGomed/auth/build/.env .