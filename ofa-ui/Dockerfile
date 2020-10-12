FROM alpine as npm

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
  && apk add --no-cache nodejs npm \
  && npm install -g cnpm --registry=https://registry.npm.taobao.org


FROM npm as builder

WORKDIR /usr/ofa
COPY . .
RUN cnpm install && npm run build


FROM nginx:stable-alpine

#添加apk源、创建用户、设置工作目录
RUN addgroup -S ofa-web && adduser ofa-web -D -G ofa-web
WORKDIR /usr/ofa
#拷贝配置文件
COPY ./ops/nginx/nginx.conf /etc/nginx/
COPY ./ops/nginx/conf.d /etc/nginx/conf.d
COPY --from=builder /usr/ofa/dist .
