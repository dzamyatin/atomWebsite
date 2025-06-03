.PHONY: app

include .env

build:
	docker compose -f docker-compose.yaml build
up:
	docker compose -f docker-compose.yaml up -d
down:
	 docker compose -f docker-compose.yaml down
dev:
	sh -c "cd frontend && yarn dev"
gen:
	go generate --tags wireinject ./...
migration-up:
	go run ./ migration-up --config config-docker.yaml
migration-down:
	go run ./ migration-down --config config-docker.yaml
migration-create:
	go run ./ migration-create --config config-docker.yaml --name new --type sql
tidy:
	sh -c "cd app && go mod tidy"
grpc:
	sh -c "cd app && go run ./ grpc --config config-local.yaml"
dev:
	sh -c "cd app && go generate ./..."
	sh -c "cd app && go generate --tags wireinject ./..."
	sh -c "cd frontend && npm run client"
	sh -c "cd frontend && npm run dev"