FROM alpine

COPY ./build/libraryservice_unix /libraryservice
COPY ./config/config-docker.yml /config.yml

RUN chmod +x /libraryservice

ENV GO_ARGS="-version=v1"

CMD ["/libraryservice", "-config", "/config.yml", "$GO_ARGS"]