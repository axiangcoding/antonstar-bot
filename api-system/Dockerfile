FROM golang:1.17-alpine as build
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o application .

FROM scratch as prod
COPY --from=build /build/application /app/application
# copy config file
COPY --from=build /build/config /app/config
EXPOSE 8888
WORKDIR /app
CMD ["/app/application"]