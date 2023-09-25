DSN := postgres://user:password@127.0.0.1:5434/cats?sslmode=disable

all: run-db install-goose | sleep migrate run-lab

sleep:
	sleep 10

install-goose:
	go get github.com/pressly/goose/cmd/goose
	go install github.com/pressly/goose/cmd/goose

migrate:
	goose --dir="deployments/migrations" postgres "$(DSN)" up

run-db:
	docker-compose --project-directory deployments up -d

run-lab:
	go run cmd/lab_06/main.go

