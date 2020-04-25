#!/bin/bash

docker run --rm \
    -v "$(pwd)":/go/src/github.com/ishansd94/reverse-proxy \
    -v "$(pwd)"/configs/proxy.yaml:/var/run/proxy/conf.yaml \
    --env-file="$(pwd)"/build/.env \
    --expose 8085 \
    -p 8085:8085 \
    -e PROJECT="github.com/ishansd94/gateway-proxy" \
    golang