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
gen:
	docker exec -ti $(PROJECTNAME)-tool bash -c "go generate --tags wireinject ./..."