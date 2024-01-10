build:
	@echo "[build] building service"
	@echo "	-> building users service..."
	@go build -o bin/users-service cmd/main.go
	@echo "	Done ✔︎"

clean:
	@echo "[clean] cleaning bin folder..."
	@rm -r bin/
	@echo "	Done ✔︎"

test:
	@echo "[test] running local tests..."
	@go test -v -count=1 -failfast ./...

.PHONY: run

run:
	@echo "[run] running service in local environment"
	@export $$(cat .env) \
		&& go run cmd/main.go