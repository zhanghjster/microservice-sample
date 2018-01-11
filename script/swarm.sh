#!/usr/bin/env bash

export MIRROR="https://0jdubf1j.mirror.aliyuncs.com"

docker-machine create -d virtualbox \
    --engine-registry-mirror ${MIRROR} \
    swarm-manager

for i in 1 2; do
    docker-machine create -d virtualbox \
        --engine-registry-mirror ${MIRROR} \
        swarm-worker-${i}
done

eval $(docker-machine env swarm-manager)

MANAGER_IP=$(docker-machine ip swarm-manager)

docker swarm init --advertise-addr ${MANAGER_IP}

TOKEN=$(docker swarm join-token -q worker)

for i in 1 2; do
    eval $(docker-machine env swarm-worker-${i})
    docker swarm join --token ${TOKEN} \
        --advertise-addr $(docker-machine ip swarm-worker-${i}) \
        ${MANAGER_IP}:2377
done