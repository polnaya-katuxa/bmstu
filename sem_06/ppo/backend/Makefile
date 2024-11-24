SERVICE_NAME = postby
MIGRATION_PG = add_ro_role
MIGRATION_MONGO = add_limits_and_reaction_types
MAIN_POSTBY = "cmd/postby/main.go"
MIGRATIONS_PG_DIR = "deployments/migrations/migrations_postgres"
MIGRATIONS_MONGO_DIR = "deployments/migrations/migrations_mongo"
#DSN = postgres://$(SERVICE_NAME):xxmU112xcjwhHQYTlPF3TFHb78YZTWQ2@90.156.209.107:5432/$(SERVICE_NAME)?sslmode=disable
#DSN = postgres://$(SERVICE_NAME):password@127.0.0.1:5435/$(SERVICE_NAME)?sslmode=disable
DSN = postgres://$(SERVICE_NAME):password@127.0.0.1:5433/$(SERVICE_NAME)?sslmode=disable

install-goose:
	go get github.com/pressly/goose/cmd/goose
	go install github.com/pressly/goose/cmd/goose

install-migrate-mongo:
	npm install -g migrate-mongo

install-goimports:
	go install golang.org/x/tools/cmd/goimports@latest

goimports:
	goimports -w .

gen-server_v1: _gen-server_v1 goimports

_gen-server_v1:
	docker run --rm \
      -v ${PWD}:/backend openapitools/openapi-generator-cli generate \
      -i /backend/api/rest/v1/swagger.yaml \
      -g go-server \
      -o /backend/internal/server/v1 \
      --additional-properties=outputAsLibrary=true,sourceFolder=.

gen-server_v2: _gen-server_v2 goimports

_gen-server_v2:
	docker run --rm \
      -v ${PWD}:/backend openapitools/openapi-generator-cli generate \
      -i /backend/api/rest/v2/swagger.yaml \
      -g go-server \
      -o /backend/internal/server/v2 \
      --additional-properties=outputAsLibrary=true,sourceFolder=.

init-mongo-migration:
	cd $(MIGRATIONS_MONGO_DIR) && migrate-mongo init

create-mongo-migration:
	cd $(MIGRATIONS_MONGO_DIR) && migrate-mongo create $(MIGRATION_MONGO)

migrate-mongo:
	cd $(MIGRATIONS_MONGO_DIR) && migrate-mongo up

migrate-down-mongo:
	cd $(MIGRATIONS_MONGO_DIR) && migrate-mongo down

create-pg-migration:
	goose --dir=$(MIGRATIONS_PG_DIR) create $(MIGRATION_PG) sql

migrate-pg:
	goose --dir=$(MIGRATIONS_PG_DIR) postgres "$(DSN)" up

migrate-down-pg:
	goose --dir=$(MIGRATIONS_PG_DIR) postgres "$(DSN)" down

run-db:
	docker-compose --project-directory deployments up -d

run-nginx:
	docker-compose --project-directory ../nginx down
	docker-compose --project-directory ../nginx up -d

run-test:
	go test ./... -cover

run-web:
	go run cmd/$(SERVICE_NAME)/main.go -cfg configs/http_config.yml -ui w

run-tech:
	go run cmd/$(SERVICE_NAME)/main.go -cfg configs/tech_config.yml -ui t

build-postby:
	mkdir -p bin
	go build -o bin/postby $(MAIN_POSTBY)

install-golint:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.52.2

run-lint:
	golangci-lint run -E dupl,errname,goconst,gofumpt,goimports,wrapcheck \
	--skip-dirs internal/server,internal/mocks,internal/containers,internal/openapi,internal/repository --timeout 5m0s

fmt:
	gofumpt -l -w .
