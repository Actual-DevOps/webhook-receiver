.PHONY: all push_image build_image

APP=webhook-receiver
GROUP=actual-devops
VERSION=$(shell cat version)-1
DOCKER_REGISTRY=ghcr.io
GOLANG_VERSION=1.22.2

BUILD_CMD='GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -ldflags="-s -w" -o webhook-receiver'

all:
	@echo 'DEFAULT:                                                         '
	@echo '   make build_image                                              '
	@echo '   make build                                                    '
	@echo '   make push_image                                               '

lint:
	golangci-lint run -v

build:
	go mod download
	$(shell echo $(BUILD_CMD))

build_image:
	@echo 'Build Docker'
	docker buildx build --build-arg GOLANG_VERSION=$(GOLANG_VERSION) \
						--build-arg BUILD_CMD=$(BUILD_CMD) \
						--platform linux/amd64 \
						-t $(DOCKER_REGISTRY)/$(GROUP)/$(APP):$(VERSION) .
	docker tag $(DOCKER_REGISTRY)/$(GROUP)/$(APP):$(VERSION) $(DOCKER_REGISTRY)/$(GROUP)/$(APP):latest

push_image:
	docker push $(DOCKER_REGISTRY)/$(GROUP)/$(APP):$(VERSION)
	docker push $(DOCKER_REGISTRY)/$(GROUP)/$(APP):latest
