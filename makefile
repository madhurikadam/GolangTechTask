#protoc -I=api --go_out=../.. --go-grpc_out=require_unimplemented_servers=false:../.. service.proto
GO_SRC_EXT = .go
GO_SRC_FILES = $(shell find . -name \*${GO_SRC_EXT})
BINARY_NAME = .builds/$(SERVICE_NAME)
SERVICE_NAME = golangtechtask
TAG=latest
DOCKER_REPO=madhuri/golangtechtask
COMPOSE_DEPS_FILES=docker-compose.yml


$(BINARY_NAME): $(GO_SRC_FILES)
	@echo "\033[32m-- Building application binary for $(SERVICE_NAME)\033[0m"

	CGO_ENABLED=0 GOOS=linux go build -v -installsuffix nocgo -o $(BINARY_NAME) cmd/main.go
.PHONY: binary
binary: $(BINARY_NAME)  ## Build application binary


docker-build: $(BINARY_NAME)  ## Build the docker image for the service
	@echo "\033[32m-- Building docker container for $(SERVICE_NAME)\033[0m"
	docker build . -t ${DOCKER_REPO}:${TAG}

docker-run:  ## Starts the containers in the compose file
	docker-compose -f $(COMPOSE_DEPS_FILES) up -d

docker-down:  ## Stop the containers in the compose file
	docker-compose -f $(COMPOSE_DEPS_FILES) down
