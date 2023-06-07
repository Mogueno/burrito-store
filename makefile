# Run tests
test:
	go test ./...
fmt:
	go fmt ./...
dev:
	docker compose up --remove-orphans
.PHONY: test dev fmt run
