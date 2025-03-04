install:
	go mod download

run: install
	set -a && . $$(pwd)/.env && set +a &&\
	go run cmd/server/main.go

.PHONY: install run