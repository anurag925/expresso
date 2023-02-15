# syntax=docker/dockerfile:1

FROM golang:1.19-alpine

RUN apk add make
RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
# This container exposes port 1323 to the outside world
EXPOSE 1323

CMD ["make", "deploy"]