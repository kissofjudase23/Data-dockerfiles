
include .env

.PHONY: build build_nc up down start restart stop tty mysql

service_name=test-sqlalchemy
compose_file=docker-compose.yml

build:
	env DOCKER_BUILDKIT=1 docker-compose -f $(compose_file) -p $(service_name) build

build_nc:
	env DOCKER_BUILDKIT=1 docker-compose -f $(compose_file) -p $(service_name) build --no-cache

up:
	docker-compose -f $(compose_file) -p $(service_name) up -d

down:
	docker-compose -f $(compose_file) -p $(service_name) down

start:
	docker-compose -f $(compose_file) -p $(service_name) down

restart:
	docker-compose -f $(compose_file) -p $(service_name) restart

stop:
	docker-compose -f $(compose_file) -p $(service_name) down

mysql:
	mysql -h 127.0.0.1 -P ${DB_PORT} -u${DB_ROOT_USER} -p${DB_ROOT_PASSWORD}

tty:
	docker exec -it $(service_name) /bin/bash
