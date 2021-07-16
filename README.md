[![codecov](https://codecov.io/gh/roman-wb/crud-products/branch/master/graph/badge.svg?token=U7EV2HFYUA)](https://codecov.io/gh/roman-wb/crud-products)
[![Build Status](https://www.travis-ci.com/roman-wb/crud-products.svg?branch=master)](https://www.travis-ci.com/roman-wb/crud-products)
![Go Report](https://goreportcard.com/badge/github.com/roman-wb/crud-products)
![Repository Top Language](https://img.shields.io/github/languages/top/roman-wb/crud-products)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/roman-wb/crud-products)
![Github Repository Size](https://img.shields.io/github/repo-size/roman-wb/crud-products)
![Lines of code](https://img.shields.io/tokei/lines/github/roman-wb/crud-products)
![License](https://img.shields.io/badge/license-MIT-green)
![GitHub last commit](https://img.shields.io/github/last-commit/roman-wb/crud-products)
![Simply the best ;)](https://img.shields.io/badge/simply-the%20best%20%3B%29-orange)

# Test project - CRUD Products

## Feature Overview
- REST API
- PostgreSQL
- Migrations
- Logger
- Tests
- Docker
- GolangCI-lint
- Development with Hot-Reload

## Get Started

### Development
```bash
cp .env.sample .env
make docker-dev-up # default listen on localhost:8080
```

### Test
```bash
make docker-test-up
make test
```

## API
| Method | URL | Description
|-|-|-|
|GET|/products|Return all products|
|POST|/products|Create new product (use JSON body)|
|GET|/products/{id}|Get product by id|
|POST|/products/{id}|Update product by id (use JSON body)|
|DELETE|/products/{id}|Delete product by id|

## Makefile commands
- `make server-run` - Run server with .env config
- `make docker-dev-up` - Run development environment and hot reload server
- `make docker-test-up` - Run test environment
- `make create-migration name={your_name}` - Create migration in dir `/migrations`
- `make generate` - Generate mocks interfaces
- `make code` - Run `code-style && code-lint`
- `make code-style` - Run `goimports`
- `make code-lint` - Run `golangci-lint`
- `make test` - Run unit + integrations tests (with docker)
- `make unit-test` - Run unit tests
- `make cover` - Open test `coverage.out` (use after `make test`)
- `make psql` - Open postgres command line with development db

## Packages
- github.com/joho/godotenv
- github.com/gorilla/mux
- github.com/jackc/pgx/v4 
- go.uber.org/zap 
- github.com/stretchr/testify
- github.com/golang/mock

## Todo
- Cache
- Auth
- Tests with Postman
- Metrics (go / pgx)

## License

MIT