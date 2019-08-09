FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/reviews-linux-amd64 app
ADD ./configs/reviews.yml .
ADD ./scripts/wait-for .
