
container_name = redis
redis_host = localhost
redis_port = 6379


.PHONY: up down tty ps
up:
	docker-compose -p $(container_name) up -d

down:
	docker-compose -p $(container_name) down


tty:
	docker exec -it $(container_name) sh

ps:
	docker-compose ps

redis:
	redis-cli -h $(redis_host) -p $(redis_port)
