postgres:
	docker run --name mobile_money -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:12-alpine
createdb:
	docker exec -it mobile_money createdb --username=root --owner=root mobile_money
dropdb:
	docker exec -it mobile_money dropdb mobile_money
migrateup:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/mobile_money?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:password@localhost:5432/mobile_money?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

format:
	go fmt ./...

peepdb:
	docker exec -it mobile_money psql -U root -d mobile_money

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/UnplugCharger/mobile_money/db/sqlc Store

security:
	gosec ./...


.PHONY: postgres createdb dropdb migrateup migratedown sqlc test format peepdb server mock security