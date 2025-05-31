DOCKER_COMPOSE_FILE = docker/docker-compose.yml
ENV_FILE = docker/.env

setup:
	cp docker/.env.example $(ENV_FILE); \

up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up -d

down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) down

build:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build

logs:
	docker-compose -f $(DOCKER_COMPOSE_FILE) logs -f

sh:
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec go-app sh

tidy:
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec go-app go mod tidy

wire:
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec go-app wire gen internal/infrastructure/di/wire.go

migrate:
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec go-app migrate create -ext sql -dir migrations -seq $(name)

migrate-up:
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec go-app migrate -path migrations -database "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable" up

migrate-down:
	docker-compose -f $(DOCKER_COMPOSE_FILE) exec go-app migrate -path migrations -database "postgres://postgres:postgres@postgres:5432/postgres?sslmode=disable" down