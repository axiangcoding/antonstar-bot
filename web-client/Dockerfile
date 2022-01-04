FROM node:16-alpine as build
WORKDIR /build
COPY . /build/
RUN yarn run build


FROM nginx:1.21-alpine as prod
WORKDIR /app
COPY --from=build /build/dist /app/dist
COPY --from=build /build/nginx/default.conf /etc/nginx/conf.d/default.conf
