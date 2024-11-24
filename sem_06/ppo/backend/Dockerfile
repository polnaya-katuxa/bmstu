FROM golang:1.20-bullseye AS builder

RUN \
  apt-get update \
  && apt-get install ca-certificates \
  && update-ca-certificates

WORKDIR /opt
COPY . /opt
ARG GOARCH=amd64
RUN go mod download all && go mod tidy && go env
RUN make build-postby

FROM amd64/debian:bullseye-slim
WORKDIR /opt
COPY --from=builder /opt/bin/postby /opt/postby
#COPY ./configs/config.yml /opt/config.yml
COPY --from=builder /etc/ssl/certs /etc/ssl/certs

EXPOSE 80

ENTRYPOINT ["/opt/postby", "-cfg", "/etc/postby/config.yml", "-ui", "w"]