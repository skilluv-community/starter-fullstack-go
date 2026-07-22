.PHONY: dev test lint build fmt clean help

help:
	@echo "Targets: dev / test / lint / build / fmt / clean"

dev:
	docker compose up --build

test:
	cd backend && go test ./...
	cd frontend && npm run test:unit -- --run

lint:
	cd backend && go vet ./... && gofmt -l .
	cd frontend && npm run lint && npm run check

build:
	cd backend && go build -o bin/server ./cmd/server
	cd frontend && npm run build

fmt:
	cd backend && gofmt -w .
	cd frontend && npm run format

clean:
	rm -rf backend/bin backend/tmp frontend/node_modules frontend/build frontend/.svelte-kit
