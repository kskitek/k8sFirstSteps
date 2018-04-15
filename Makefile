# .PHONY

all: deps test run

PROJECT_NAME=k8sFirstSteps
PROJECT_REPOSITORY=github.com/KSkitek/$(PROJECT_NAME)
DOCKER_REGISTRY=none_yet

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
	env GOOS=linux go build -o $(PROJECT_NAME)-linux

dbuild: compile-linux
	docker build -t $(DOCKER_REGISTRY) .

drun: dbuild
	docker run --rm -it -p 8080:8080 $(DOCKER_REGISTRY)

dpush: dbuild
	docker push $(DOCKER_REGISTRY)