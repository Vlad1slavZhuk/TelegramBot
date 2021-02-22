run:
	go run cmd/bot/main.go

run-db:
	docker-compose up

exec-db:
	docker exec -it postgres psql -U postgres