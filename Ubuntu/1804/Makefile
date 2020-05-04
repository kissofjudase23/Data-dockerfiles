
container_name = ubuntu1804_env

.PHONY: up
up:
	docker-compose -p $(container_name) up -d --build


.PHONY: down
down:
	docker-compose down


.PHONY: tty
tty:
	docker exec -it $(container_name) /bin/bash


.PHONY: ps
ps:
	docker-compose ps
