include .env

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=nps
BINARY_UNIX=$(BINARY_NAME)_unix

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
	$(GOBUILD) -o $(BINARY_NAME) -v

test: 
	$(GOTEST) -v ./...

clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -f $(BINARY_UNIX)

run: 
	./$(BINARY_NAME) -env production

deploy: build run

deploy_staging:build
	./$(BINARY_NAME) -env staging

debug:
	$(GOCMD) run .

deps:
	$(GOGET) github.com/markbates/goth
	$(GOGET) github.com/markbates/pop

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


# Cross compilation
build-linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_UNIX) -v
docker-build:
	docker run --rm -it -v "$(GOPATH)":/go -w /go/src/bitbucket.org/rsohlich/makepost golang:latest go build -o "$(BINARY_UNIX)" -v

