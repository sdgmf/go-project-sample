FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/ratings-linux-amd64 app
ADD ./configs/ratings.yml .
ADD ./scripts/wait-for .


