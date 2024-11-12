#!/bin/bash

echo "Stopping containers and removing associated volumes"
docker-compose down -v
echo "Building containers"
docker-compose up --build
