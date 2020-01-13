include .env
export

include Makefile.common

COMPOSE_FILES := -f docker-compose.yml

all: | build up

.PHONY: build
build:
	docker-compose -p ${PROJECT_NAME} ${COMPOSE_FILES} build --pull --force-rm --no-cache

.PHONY: up
up: | docker-up

.PHONY: down
down: | docker-down

.PHONY: restart
restart: | docker-restart

.PHONY: docker-up
docker-up:
	docker-compose -p ${PROJECT_NAME} ${COMPOSE_FILES} up -d --force-recreate

.PHONY: docker-down
docker-down:
	docker-compose -p ${PROJECT_NAME} ${COMPOSE_FILES} down -t 3

.PHONY: docker-restart
docker-restart:
	docker-compose -p ${PROJECT_NAME} ${COMPOSE_FILES} restart -t 1

.PHONY: sh
sh:
	docker-compose -p ${PROJECT_NAME} ${COMPOSE_FILES} exec app sh

.PHONY: bash
bash:
	docker-compose -p ${PROJECT_NAME} ${COMPOSE_FILES} exec app bash

.PHONY: app-build
app-build:
	docker-compose -p ${PROJECT_NAME} ${COMPOSE_FILES} exec -T app make -f Makefile.native build
