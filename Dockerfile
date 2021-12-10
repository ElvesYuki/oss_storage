FROM debian:bullseye-slim

WORKDIR /go/bin
COPY ./oss_storage ./oss_storage
COPY ./config.yaml ./config.yaml

EXPOSE 9091

CMD ["./oss_storage"]