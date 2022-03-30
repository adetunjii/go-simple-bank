migrate_up:
	migrate -path db/migration -database "postgresql://teej4y:password@localhost:5432/simplebank?sslmode=disable" -verbose up

migrate_down:
	migrate -path db/migration -database "postgresql://teej4y:password@localhost:5432/simplebank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen --build_flags=--mod=mod --package mockdb --destination db/mock/store.go github.com/Adetunjii/simplebank/db/repository IStore