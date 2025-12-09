all: build

build:
	go build -o bin/manager main.go

docker-build:
	docker build -t feroz29/llm-log-operator:latest .

docker-push:
	docker push feroz29/llm-log-operator:latest
