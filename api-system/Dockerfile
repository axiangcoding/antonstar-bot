FROM golang:1.20-alpine as build
#如果网络访问慢，可以启用以下配置项
#ENV GOPROXY=https://goproxy.cn,direct
WORKDIR /build
COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o application ./cmd/app

FROM alpine:latest as prod
LABEL maintainer=axiangcoding@gmail.com
ARG VERSION
WORKDIR /app
ENV AS_APP_VERSION=${VERSION}
COPY --from=build /build/application /app/application
# 复制config文件夹到镜像中
COPY --from=build /build/config /app/config
EXPOSE 8888
CMD ["/app/application"]