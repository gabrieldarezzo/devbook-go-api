start:
	go run main.go
up:
	docker compose up -d

up-force:
	docker compose up --build --force-recreate