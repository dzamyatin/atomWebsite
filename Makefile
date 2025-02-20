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
protoc:
	docker exec -ti $(PROJECTNAME)-tool bash -c "cd /app/proto && /home/$(USERNAME)/protoc/bin/protoc --go_out=../internal/grpc/generated/ --go_opt=paths=source_relative \
                 --go-grpc_out=../internal/grpc/generated/ --go-grpc_opt=paths=source_relative \
                 *.proto"