FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct GOSUMDB="sum.golang.google.cn" \
    go build -v -o docker_test/build/main docker_test/main.go
