#!/bin/bash

echo Encerrando o ambiente local

docker compose -f ./mysql/docker-compose.yml down

docker volume prune