
include .env

.PHONY: build build_nc up down start restart stop tty mysql

service_name=test-gorm
compose_file=docker-compose.yml

build:
	env DOCKER_BUILDKIT=1 docker-compose -f $(compose_file) -p $(service_name) build

build_nc:
	env DOCKER_BUILDKIT=1 docker-compose -f $(compose_file) -p $(service_name) build --no-cache

up:
	docker-compose -f $(compose_file) -p $(service_name) up -d

down:
	docker-compose -f $(compose_file) -p $(service_name) down

downv:
	docker-compose -f $(compose_file) -p $(service_name) down --volumes

start:
	docker-compose -f $(compose_file) -p $(service_name) start

restart:
	docker-compose -f $(compose_file) -p $(service_name) restart

stop:
	docker-compose -f $(compose_file) -p $(service_name) stop

mysql:
	mycli -h 127.0.0.1 -P ${DB_PORT} -u${DB_ROOT_USER} -pq

login:
	docker exec -it $(service_name) /bin/bash

build_go:
	go build ./test_gorm.go

run_go:
	go run ./test_gorm.go
