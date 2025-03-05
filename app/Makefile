docker:
	docker build --tag cds-superset-docs --platform linux/arm64 .

fmt:
	go fmt ./...

install:
	go mod download

lint:
	golangci-lint run ./...

run: install
	set -a && . $$(pwd)/.env && set +a &&\
	go run cmd/server/main.go

test:
	go test -v ./...

.PHONY: docker fmt install lint run test