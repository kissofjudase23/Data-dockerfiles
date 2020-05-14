
container_name = prom

.PHONY: up
up:
	docker-compose -p $(container_name) up -d

.PHONY: down
down:
	docker-compose -p $(container_name) down

.PHONY: tty
tty:
	docker exec -it $(container_name) sh


.PHONY: ps
ps:
	docker-compose ps