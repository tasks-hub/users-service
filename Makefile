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
	@docker-compose up --build

CONTAINER_NAME = users-service-goose-1

run-migrations:
	docker exec -it $(CONTAINER_NAME) chmod +x ./run_migrations.sh
	docker exec -it $(CONTAINER_NAME) ./run_migrations.sh
