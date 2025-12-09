all: build

build:
	go build -o bin/manager main.go

docker-build:
	docker build -t yourrepo/llm-log-operator:latest .

docker-push:
	docker push yourrepo/llm-log-operator:latest
