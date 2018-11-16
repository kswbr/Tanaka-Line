FROM alpine:latest

RUN apk update && apk upgrade && apk add --update --no-cache alpine-sdk tzdata nodejs protobuf npm
RUN apk add --update --no-cache \
      bash \
      git \
      openssh \
      yaml

RUN mkdir /app
WORKDIR /app

ADD package.json .
ADD index.js .
ADD proto .

EXPOSE 8081
