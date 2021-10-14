#!/bin/sh

docker network create tpm-network

docker run -d -p 5432:5432 --net tpm-network --name pg \
  -e POSTGRES_DB=tpm \
  -e POSTGRES_USER=tpm \
  -e POSTGRES_PASSWORD=tpm \
  postgres

docker build -t tpm .

docker run -d -p 9000:9000 --net tpm-network --name tpm tpm

