
container_name = centos8_dev

.PHONY: up
up:
	docker-compose -p $(container_name) up -d --build

.PHONY: down
down:
	docker-compose down

.PHONY: tty
tty:
	docker exec -it $(container_name) /bin/bash
