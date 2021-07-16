include .env

.PHONY: server-run docker-dev-up create-migration generate code code-style code-lint docker-test-up test unit-test cover pre-commit psql

# Development

server-run:
	go run cli/server/main.go

docker-dev-up:
	docker-compose -f docker-compose.dev.yml up

create-migration:
	docker run --rm -it -v `pwd`/migrations:/migrations --network host migrate/migrate create -ext sql -dir=/migrations $(name)

generate:
	go generate ./...

# CI/CD

code: code-style code-lint

code-style:
	goimports -w ./..

code-lint:
	docker run --rm -v `pwd`:/app -w /app golangci/golangci-lint golangci-lint run -v

docker-test-up:
	docker-compose -p crud-products_test -f docker-compose.test.yml up --remove-orphans --build

test:
	go test -count 1 -race -coverprofile=coverage.out ./...

unit-test:
	go test -short -count 1 -race -coverprofile=coverage.out ./...

cover:
	go tool cover -func -html=coverage.out

pre-commit:
	make code
	make test

# Utils

psql:
	psql ${DATABASE_URL}