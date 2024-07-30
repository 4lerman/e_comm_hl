build:
	go build -o bin/ecom api/main.go
test:
	go test -v ./...
run: build
	./bin/ecom

migration:
	migrate create -ext sql -dir cmd/migrate/migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	go run common/migrate/main.go up

migrate-down:
	go run common/migrate/main.go down