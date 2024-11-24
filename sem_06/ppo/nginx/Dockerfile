FROM golang:1.19
COPY ./backend /backend
RUN cd /backend && mkdir -p bin && go build -o /backend/bin/postby /backend/cmd/postby/main.go
ENTRYPOINT ["/backend/bin/postby", "-cfg", "/backend/configs/http_config.yml", "-ui", "w"]