
container_name = ubuntu_1804_test

.PHONY: up
up:
	docker-compose up -d --build

.PHONY: down
down:
	docker-compose down

.PHONY: tty
tty:
	docker exec -it $(container_name) /bin/bash

