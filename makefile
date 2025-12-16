# =========================
# Project config
# =========================
APP_NAME := ginject-cms-api
CMD_DIR  := ./cmd/app
BIN_DIR  := bin
BIN_FILE := ${BIN_DIR}/${APP_NAME}

GO       := go
GOFLAGS  := -v

DOCKER := docker
COMPOSE := ${DOCKER} compose

COMPOSE_INFRA_FILE := deploy/docker/docker-compose.infra.yml
COMPOSE_FILE := deploy/docker/docker-compose.yml
DOCKER_FILE := deploy/docker/Dockerfile
ENV_FILE := .env 

.PHONY: help
help: ## Show help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' ${MAKEFILE_LIST} \
	| awk 'BEGIN {FS = ":.*?## "}; {printf "  \033[36m%-20s\033[0m %s\n", $$1, $$2}'

# =========================
# Go commands
# =========================
.PHONY: deps
deps: ## Download dependencies
	yarn
	go mod download

.PHONY: upgrade
upgrade: ## Upgrade dependencies
	go get -u ./...

.PHONY: tidy
tidy: ## Go mod tidy
	go mod tidy

.PHONY: run
run: ## Run app locally
	@test -f .env || cp .env.example .env
	go run cmd/app/main.go

.PHONY: run-build
run-build: ## Run binary locally
	./${BIN_FILE}

.PHONY: test
test: ## Run tests
	go test ./... -cover

.PHONY: build
build: ## Build binary
	mkdir -p ${BIN_DIR}
	CGO_ENABLED=0 GOARCH=amd64 go build ${GOFLAGS} -o ${BIN_FILE} ${CMD_DIR}

.PHONY: lint
lint: ## Run linter
	golangci-lint run

.PHONY: clean
clean: ## Clean build artifacts
	rm -rf ${BIN_DIR}

# =========================
# Docker
# =========================
.PHONY: docker-build
docker-build: ## Build application's Dockerfile
	${DOCKER} build -t ${APP_NAME}:latest -f ${DOCKER_FILE} .

# =========================
# Docker compose
# =========================
.PHONY: compose-config
compose-config: ## Show docker-compose config
	${COMPOSE} --env-file ${ENV_FILE} -f ${COMPOSE_INFRA_FILE} -f ${COMPOSE_FILE} config

.PHONY: compose-up
compose-up: ## Start docker-compose
	${COMPOSE} -p ${APP_NAME} --env-file ${ENV_FILE} -f ${COMPOSE_INFRA_FILE} -f ${COMPOSE_FILE} up -d

.PHONY: compose-down
compose-down: ## Remove docker-compose
	${COMPOSE} -p ${APP_NAME} --env-file ${ENV_FILE} -f ${COMPOSE_INFRA_FILE} -f ${COMPOSE_FILE} down -v

.PHONY: compose-logs
compose-logs: ## Show docker-compose logs
	${COMPOSE} --env-file ${ENV_FILE} -f ${COMPOSE_INFRA_FILE} -f ${COMPOSE_FILE} logs -f

.PHONY: compose-infra-config
compose-infra-config: ## Show docker-compose infrastructure config
	${COMPOSE} --env-file ${ENV_FILE} -f ${COMPOSE_INFRA_FILE} config

.PHONY: compose-infra-up
compose-infra-up: ## Start docker-compose infrastructure 
	${COMPOSE} --env-file ${ENV_FILE} -f ${COMPOSE_INFRA_FILE} up -d

.PHONY: compose-infra-down
compose-infra-down: ## Remove docker-compose infrastructure 
	${COMPOSE} --env-file ${ENV_FILE} -f ${COMPOSE_INFRA_FILE} down -v

.PHONY: compose-infra-logs
compose-infra-logs: ## Show docker-compose infrastructure logs
	${COMPOSE} --env-file ${ENV_FILE} -f ${COMPOSE_INFRA_FILE} logs -f
