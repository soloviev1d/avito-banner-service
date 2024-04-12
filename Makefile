build:
	docker-compose --env-file creds.env build
run: build
	docker-compose --env-file creds.env up -d
kill:
	docker kill $(docker ps -q)
