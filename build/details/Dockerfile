FROM alpine:latest

RUN apk update && apk add bash

ADD ./dist/details-linux-amd64 app
ADD ./configs/details.yml .
ADD ./scripts/wait-for .


