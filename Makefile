REPO ?= nadavbm/controller
TAG ?= latest

.PHONY: docker-build
docker-build:
	docker build -t ${REPO}:${TAG} .

.PHONY: docker-push
docker-push:
	docker push ${REPO}:${TAG}