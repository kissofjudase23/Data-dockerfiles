
container_name = prom

.PHONY: up
up:
	docker-compose -p $(container_name) up -d

.PHONY: down
down:
	docker-compose down

.PHONY: tty
tty:
	docker exec -it $(container_name) /bash


.PHONY: ps
ps:
	docker-compose ps