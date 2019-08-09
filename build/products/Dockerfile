FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/products-linux-amd64 app
ADD ./configs/products.yml .
ADD ./scripts/wait-for .


