#!/usr/bin/env sh

BIN="accountservice-linux-amd64"
IMAGE="zhanghjster/accountservice:dev"

GOOS=linux go build -o ${BIN}

docker build -t ${IMAGE} .
docker push ${IMAGE}

rm -f ${BIN}