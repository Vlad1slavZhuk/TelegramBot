run:
	go run cmd/main.go

run-db:
	docker-compose up

exec-db:
	docker exec -it postgres psql -U postgres