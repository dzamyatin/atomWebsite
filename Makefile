.PHONY: app

include .env

build:
	docker compose -f docker-compose.yaml build
up:
	docker compose -f docker-compose.yaml up -d
down:
	 docker compose -f docker-compose.yaml down
gen:
	sh -c "cd app &&go generate --tags wireinject ./..."
migration-up:
	sh -c "cd app &&go run ./ migration-up --config config-local.yaml"
migration-down:
	sh -c "cd app &&go run ./ migration-down --config config-local.yaml"
migration-create:
	sh -c "cd app &&go run ./ migration-create --config config-local.yaml --name new --type sql"
tidy:
	sh -c "cd app && go mod tidy"
wire:
	sh -c "cd app/internal/di && go tool wire"
grpc:
	sh -c "cd app && go run ./ grpc --config config-local.yaml"
dev:
	sh -c "cd app && go generate ./..."
	sh -c "cd app && go generate --tags wireinject ./..."
	sh -c "cd frontend && npm run client"
	sh -c "cd frontend && npm run dev"