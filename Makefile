start:
	docker compose --env-file ./server/.env  -f ./docker-compose.yaml start

stop:
	docker compose --env-file ./server/.env  -f ./docker-compose.yaml stop

down:
	docker compose --env-file ./server/.env  -f ./docker-compose.yaml down

up:
	docker compose --env-file ./server/.env  -f ./docker-compose.yaml up

up-quite:
	docker compose --env-file ./server/.env  -f ./docker-compose.yaml up -d

build:
	docker compose --env-file ./server/.env  -f ./docker-compose.yaml build

server-restart:
	docker compose --env-file ./server/.env  -f ./docker-compose.yaml restart server

db-up:
	docker compose --env-file ./server/.env  -f ./docker-compose.yaml up -d taskmanagement_db