DOCKER_COMPOSE_CMD := $(shell which docker-compose)
DOCKER_COMPOSE_MYSQL_PREFIX := $(DOCKER_COMPOSE_CMD) -f tools/docker/docker-compose.mysql.yml

run-mysql-local: 
	$(DOCKER_COMPOSE_MYSQL_PREFIX) up -d --build