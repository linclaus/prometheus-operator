FROM golang:1.13
WORKDIR /
COPY bin/main main
COPY config.yaml config.yaml

ENTRYPOINT ["/main"]