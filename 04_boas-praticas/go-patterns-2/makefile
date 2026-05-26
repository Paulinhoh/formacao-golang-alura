run:
	go run ./cmd/api/main.go

lint:
	golangci-lint run

lint-docker:
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v2.12.2 golangci-lint run

test: # teste do database
	docker compose exec api go test ./internal/config -v 

start:
	docker compose up -d

stop:
	docker compose down

ci: start test lint-docker
