# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=libraryservice
BINARY_UNIX=$(BINARY_NAME)_unix
DOCKER_BUILD_IMAGE=$(BINARY_NAME)-build
MAIN_GO_FILE=main.go
DOCKER_IMAGE_TAG=library-service

all: test build
build: 
	$(GOBUILD) -o $(BINARY_NAME) -v
test: 
	$(GOTEST) -v ./...
clean: 
	$(GOCLEAN)
	rm -f $(BINARY_NAME)
	rm -rf build
run:
	$(GOBUILD) -o $(BINARY_NAME) $(MAIN_GO_FILE)
	./$(BINARY_NAME)
deps:
	$(GOGET) github.com/gorilla/mux
	$(GOGET) github.com/prometheus/client_golang/prometheus
	$(GOGET) gopkg.in/yaml.v2
build-linux:
	mkdir -p build
    $(GOBUILD) -o "build/$(BINARY_UNIX)" $(MAIN_GO_FILE)
docker-build-deps:
	docker build -t $(DOCKER_BUILD_IMAGE) -f build.Dockerfile .
docker-build:
	mkdir -p build
	docker run --rm -v "$$PWD:/go/src/library-service" -it $(DOCKER_BUILD_IMAGE) go build -o "build/$(BINARY_UNIX)" $(MAIN_GO_FILE)
docker-build-image: docker-build
	docker build -t $(DOCKER_IMAGE_TAG) .
    