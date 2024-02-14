DOCKER_COMPOSE = docker-compose
DOCKER_COMPOSE_FILE = docker-compose.yml
CONTAINER_NAME = tasks-hub-users
GOOSE_CONTAINER_NAME = ${CONTAINER_NAME}-goose-1

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
	@go test -v -count=1 -cover -failfast ./...

run-migrations:
	docker exec -it $(GOOSE_CONTAINER_NAME) chmod +x ./run_migrations.sh
	docker exec -it $(GOOSE_CONTAINER_NAME) ./run_migrations.sh

run-local:
	@echo "[run-local] running service in local environment"
	@export $$(cat .env) \
		&& go run cmd/main.go

# Run the service in a local environment using Docker Compose
run-docker:
	@echo "[run-docker] running service in local environment"
	@$(DOCKER_COMPOSE) up --build

# Stop and remove the Docker Compose services
down:
	@echo "[down] stopping and removing services"
	@$(DOCKER_COMPOSE) down