#编译
FROM golang:alpine AS builder
MAINTAINER gedongdong@extrotec.com
WORKDIR /go
COPY . /go
ENV GOPROXY https://goproxy.cn,direct
ENV GO111MODULE on
RUN unset GOPATH && GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-w -s"  .

#运行
FROM alpine AS runner
# 全局工作目录
WORKDIR /go
COPY --from=builder /go/go-exporter .
RUN echo "https://mirrors.aliyun.com/alpine/v3.8/main/" > /etc/apk/repositories \
    && echo "https://mirrors.aliyun.com/alpine/v3.8/community/" >> /etc/apk/repositories \
    && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime  \
    && echo Asia/Shanghai > /etc/timezone \
    && apk del tzdata
EXPOSE 8080
ENTRYPOINT ["./go-exporter"]
