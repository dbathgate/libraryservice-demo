FROM golang:1.10-alpine as builder

RUN apk --update upgrade \
 && apk --no-cache --no-progress add make git \
 && rm -rf /var/cache/apk/*

WORKDIR /go/src/library-service
COPY . /go/src/library-service

RUN ["make", "deps"]
RUN ["make", "build-linux"]

FROM alpine

COPY --from=builder /go/src/library-service/libraryservice_unix /libraryservice
COPY --from=builder /go/src/library-service/config/config-docker.yml /config.yml
COPY --from=builder /go/src/library-service/run.sh /run.sh

RUN chmod +x /libraryservice \
 && chmod +x /run.sh

ENV GO_ARGS="-version=v1"

CMD ["/run.sh"]