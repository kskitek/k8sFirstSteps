# .PHONY

all: deps test run

PROJECT_NAME=k8sFirstSteps
VERSION=0.1
PROJECT_REPOSITORY=github.com/KSkitek/$(PROJECT_NAME)
DOCKER_REGISTRY=eu.gcr.io/jugpoznan2018/k8sfirststeps:$(VERSION)

deps:
	dep ensure

test:
	go test -race ./...

verify:
	go fmt ./...
	go vet ./...
	go test ./... -race
	staticcheck .
	gosimple .

compile: test
	go build

run: compile
	./$(PROJECT_NAME)

compile-linux: test
	env GOOS=linux CGO_ENABLED=0 go build -o $(PROJECT_NAME)-linux

dbuild: compile-linux
	docker build -t $(DOCKER_REGISTRY) -f Dockerfile-single .

drun: dbuild
	docker run --rm -it -p 8080:8080 $(DOCKER_REGISTRY)

dpush: dbuild
	@ docker push $(DOCKER_REGISTRY)
	@# gcloud auth configure-docker
	@# gcloud docker --authorize-only
	@#gcloud docker -- push $(DOCKER_REGISTRY)
	@# 
	@ echo "Docker image: $(DOCKER_REGISTRY)"