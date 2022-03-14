FROM golang:1.17-alpine as build
ARG AS_VERSION
ENV GOPROXY https://goproxy.cn,direct
WORKDIR /build
COPY . /build/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o application .
RUN echo "${AS_VERSION}"

FROM busybox:stable as prod
ENV AS_API_APP_VERSION ${AS_VERSION}
RUN echo "${AS_VERSION}"
RUN echo "${AS_API_APP_VERSION}"
COPY --from=build /build/application /app/application
# copy config file
COPY --from=build /build/config /app/config
EXPOSE 8888
WORKDIR /app
CMD ["/app/application"]