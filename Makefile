DB_URL=postgresql://postgres:12345678@localhost:5433/postgres?sslmode=disable
mup:
	goose postgres "$(DB_URL)" up
mdown:
	goose postgres "$(DB_URL)" down
new_migration:
	migrate create -ext sql -dir app/db/migration -seq $(name)
mforce:
	migrate -path app/db/migration -database "$(DB_URL)" -verbose force 1
migrateup-github:
	migrate -path app/db/migration -database "$(DB_URL)" -verbose up
	 
# sqlc:
# 	docker run --rm -v ".://src" -w //src sqlc/sqlc:1.20.0 generate 

sqlc:
	sqlc generate

postgres:
	docker run -d  --name postgres  -p 5433:5432 -e POSTGRES_PASSWORD=12345678  -e PGDATA=/var/lib/postgresql/data/pgdata  -v postgres_volume:/var/lib/postgresql/data  postgres:15-alpine

redis:
	docker run -d --name redis -p 6379:6379 redis:7-alpine

rabbitmq:
	docker run -d --name rabbitmq -p 5672:5672 -p 15672:15672 rabbitmq:3-management

.PHONY: mup mdown new_migration mforce sqlc postgres migrateup-github redis mup_test mdown_test mforce_test rabbitmq