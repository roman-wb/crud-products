include .env

.PHONY: server-run docker-dev-up create-migration code code-style code-lint docker-test-up docker-test-up generate test unit-test cover psql

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

test:
	docker-compose -f docker-compose.test.yml up -d --remove-orphans --build
	./wait-for-migrate.sh "go test -count 1 -race -coverprofile=coverage.out `go list ./... | grep -v ./mock_handlers | grep -v ./test`"
	docker-compose -f docker-compose.test.yml down --volume

unit-test:
	go test -count 1 -short -race -coverprofile=coverage.out `go list ./... | grep -v ./mock_handlers | grep -v ./test`

cover:
	go tool cover -html=coverage.out

# Utils

psql:
	psql ${DATABASE_URL}