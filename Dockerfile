FROM golang:1.20-alpine as build-backend
# ENV GOPROXY=https://goproxy.cn
WORKDIR /build
COPY ./api-system /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o application ./cmd/app


FROM node:lts-alpine as build-frontend
# RUN npm config set registry https://registry.npmmirror.com
WORKDIR /build
COPY ./browser-client /build/
RUN npm install
RUN npm run build

FROM alpine:latest as prod
LABEL maintainer=axiangcoding@gmail.com
ARG VERSION
WORKDIR /app
ENV AS_APP_VERSION=${VERSION}
COPY --from=build-backend /build/application /app/application
# 复制config文件夹到镜像中
COPY --from=build-backend /build/config /app/config
COPY --from=build-frontend /build/dist /app/web
EXPOSE 8888
CMD ["/app/application"]