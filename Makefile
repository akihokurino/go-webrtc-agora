MAKEFLAGS=--no-builtin-rules --no-builtin-variables --always-make
ROOT := $(realpath $(dir $(lastword $(MAKEFILE_LIST))))

INTERNAL_TOKEN := ""

vendor:
	go mod tidy

gen:
	gqlgen
	cp di/wire_gen.default.go di/wire_gen.go
	go generate di/wire_gen.go

build:
	GOOS=linux GOARCH=amd64 go build -o .tmp/main ./entrypoint/

run-local:
	docker-compose up

deploy-gae: vendor gen build gen-app-yaml
	gcloud app deploy --quiet --version 1 --project canvas-329810 app.yaml

deploy-index:
	gcloud app deploy --quiet index.yaml

gen-app-yaml:
	go run cmd/merge_yaml/main.go app.yaml app.template.yaml env.yaml