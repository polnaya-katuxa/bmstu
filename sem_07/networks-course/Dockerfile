FROM gcc

COPY . /build
WORKDIR /build
RUN gcc *.c -o /opt/server

WORKDIR /static
ENTRYPOINT ["/opt/server", "80"]