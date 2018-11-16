FROM golang:1-alpine
RUN apk update && apk upgrade && apk add --update --no-cache alpine-sdk tzdata nodejs protobuf
RUN apk add --update --no-cache \
      bash \
      git \
      openssh \
      mysql \
      mysql-client \
      tzdata \
      yaml

RUN mkdir /go/src/app
WORKDIR /go/src/app

RUN go get -u "github.com/PuerkitoBio/goquery"
RUN go get -u "google.golang.org/grpc"

EXPOSE 8080
