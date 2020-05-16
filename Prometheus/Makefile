
container_name = prom

.PHONY: up down tty p

up:
	docker-compose -p $(container_name) up -d

down:
	docker-compose -p $(container_name) down

tty:
	docker exec -it $(container_name) sh

ps:
	docker-compose ps