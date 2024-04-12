build:
	docker-compose --env-file creds.env build
run: build
	docker-compose --env-file creds.env up -d
kill:
	docker kill $(docker ps -q)
local:
	POSTGRES_USER=postgres POSTGRES_PASSWORD=postgres123 go run main.go
