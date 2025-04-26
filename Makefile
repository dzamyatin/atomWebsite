.PHONY: app

include .env

app:
	cd app; \
	go get; \
	go build -o exec/app
build:
	docker compose -f docker-compose.yaml build
up:
	docker compose -f docker-compose.yaml up -d
down:
	 docker compose -f docker-compose.yaml down
logs:
	docker logs -f $(PROJECTNAME)-tool
sh:
	docker exec -ti $(PROJECTNAME)-tool bash
shf:
	docker exec -ti $(PROJECTNAME)-frontend sh
dev:
	docker exec -ti $(PROJECTNAME)-frontend yarn dev
gen:
	docker exec -ti $(PROJECTNAME)-tool bash -c "go generate --tags wireinject ./..."
migration-up:
	docker exec -ti $(PROJECTNAME)-tool bash -c "go run ./ migration-up --config config-docker.yaml"
migration-down:
	docker exec -ti $(PROJECTNAME)-tool bash -c "go run ./ migration-down --config config-docker.yaml"
migration-create:
	docker exec -ti $(PROJECTNAME)-tool bash -c "go run ./ migration-create --config config-docker.yaml --name new --type sql"
tidy:
	sh -c "cd app && go mod tidy"