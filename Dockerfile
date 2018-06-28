FROM alpine

COPY ./build/libraryservice_unix /libraryservice
COPY ./config/config-docker.yml /config.yml

RUN chmod +x /libraryservice

CMD ["/libraryservice", "-config", "/config.yml", "-version", "v2"]