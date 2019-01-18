SHELL := /bin/bash
DOCKER_NAMESPACE := emzian7
DOCKER_IMAGE_NAME := go-dev

start:
    @docker run --rm
        -v $(pwd):/go/src/github.com/ishansd94/go-reverse-proxy \
        -v $(pwd)/configs/proxy.yaml:/var/run/proxy/conf.yaml \
        --env-file=$(pwd)/build/.env \
        -e PROJECT="github.com/ishansd94/go-reverse-proxy" \
        $(DOCKER_NAMESPACE)/$(DOCKER_IMAGE_NAME)

.PHONY: help init pull branch pr
.DEFAULT_GOAL := help