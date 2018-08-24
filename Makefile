VERSION ?= latest

.PHONY: docker-build
docker-build:
	GOOS=linux GOARCH=amd64 go build -o bin/github-pr-commits *.go
	docker build -t mattfarina/github-pr-commits:$(VERSION) .

# You must be logged into DOCKER_REGISTRY before you can push.
.PHONY: docker-push
docker-push:
	docker push mattfarina/github-pr-commits