FROM golang:latest as builder

WORKDIR /app

COPY . .

RUN GOOS=linux GOARCH=amd64 GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct GOSUMDB="sum.golang.google.cn"
#    go build -v -o build/main main.go

# 创建容器后执行
CMD ["go",  "build", "-v", "-o", "build/main", "main.go"]
#ENTRYPOINT ["go",  "build", "-v", "-o", "build/main", "main.go"]