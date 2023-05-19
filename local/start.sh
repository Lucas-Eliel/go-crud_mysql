#!/bin/bash

echo Iniciando ambiente local

docker compose -f ./mysql/docker-compose.yml up -d 
