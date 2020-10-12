FROM golang:1.14.3-alpine3.11 as cache
LABEL stage=cache
ENV GOPROXY https://goproxy.cn,https://goproxy.io,https://mirrors.aliyun.com/goproxy/,https://athens.azurefd.net,direct
WORKDIR /usr/src/ofa


FROM cache as builder
LABEL stage=intermediate
COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ofa -tags static -ldflags '-s -w' main.go


FROM alpine:3.11.6
WORKDIR /usr/src/ofa
RUN apk --no-cache add ca-certificates tzdata && \
    ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    echo "Asia/Shanghai" > /etc/timezone
COPY --from=builder /usr/src/ofa/ofa ./
COPY --from=builder /usr/src/ofa/conf.online ./conf
EXPOSE 8086
CMD ["./ofa"]