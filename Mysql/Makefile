
container_name = mysql
mysql_host = 127.0.0.1
mysql_port = 33060
mysql_usr = tomlin
mysql_pwd = q


.PHONY: up down tty ps redis

up:
	docker-compose -p $(container_name) up -d

down:
	docker-compose -p $(container_name) down

tty:
	docker exec -it $(container_name) sh

ps:
	docker-compose ps

mysql:
	mysql -h $(mysql_host) -P $(mysql_port) -u$(mysql_usr) -p$(mysql_pwd)
