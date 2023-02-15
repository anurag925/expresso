include .env

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
SERVER_BINARY_NAME=expresso
WORKER_BINARY_NAME=expresso-worker

# Relative address
SERVER=./cmd/server/main.go
WORKER=./cmd/worker/main.go

# Migration parameters
DATABASE_URL=mysql://${DB_USERNAME}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}
MIGRATION_LOCATION=./pkg/repositories/db/migrations
GO_MIGRATE=migrate -verbose
MIGRATE=${GO_MIGRATE} -path ${MIGRATION_LOCATION} -database "${DATABASE_URL}"

# SqlBoiler
SQLBOILER=sqlboiler
DB_CONFIG_FILE=./pkg/repositories/db/sqlboiler.toml
DB_CONFIG_TO_USE=mysql


all: debug

build: 
	$(GOBUILD) -o $(SERVER_BINARY_NAME) -v ${SERVER}

build_worker: 
	$(GOBUILD) -o $(WORKER_BINARY_NAME) -v ${WORKER}

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(SERVER_BINARY_NAME)
	rm -f $(WORKER_BINARY_NAME)

run: 
	./$(SERVER_BINARY_NAME)

run_worker:
	./$(WORKER_BINARY_NAME)

server: build run
worker: build_worker run_worker

debug:
	$(GOCMD) run ${SERVER}

# Migration
create_migration:
	${GO_MIGRATE} create -ext sql -dir ${MIGRATION_LOCATION} -seq ${name}

migrate_up:
	${MIGRATE} up

migrate_up_to:
	${MIGRATE} up ${version}

migrate_down_all:
	${MIGRATE} down

migrate_down:
	${MIGRATE} down 1

migrate_down_by:
	${MIGRATE} down ${versions}

migrate_version:
	${MIGRATE} version

migrate_drop:
	${MIGRATE} drop

migrate_force:
	${MIGRATE} force ${version}

models_update:
	${SQLBOILER} ${DB_CONFIG_TO_USE} -c ${DB_CONFIG_FILE}

