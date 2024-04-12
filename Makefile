build:
	docker-compose --env-file creds.env build
all: build
	docker-compose --env-file creds.env up -d
