DB_URL=postgresql://postgres:167916@localhost:5432/simple-bank?sslmode=disable

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

new_migration:
	migrate create -ext sql -dir db/migration -digits 3 -seq $(name)

.PHONY: migrateup migratedown new_migration