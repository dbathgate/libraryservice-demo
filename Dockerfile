FROM alpine

COPY ./build/libraryservice_unix /libraryservice
COPY ./config/config-docker.yml /config.yml
COPY ./run.sh /run.sh

RUN chmod +x /libraryservice \
 && chmod +x /run.sh

ENV GO_ARGS="-version=v1"

CMD ["/run.sh"]