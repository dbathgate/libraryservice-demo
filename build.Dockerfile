FROM golang:1.10-alpine

RUN apk --update upgrade \
 && apk --no-cache --no-progress add make git \
 && rm -rf /var/cache/apk/*

WORKDIR /go/src/library-service
COPY . /go/src/library-service

RUN ["make", "deps"]