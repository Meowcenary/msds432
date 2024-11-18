General Architecture
---
This project is divided into several microservices that are packaged with Docker
. Instructions for building and running the individual microservices can be
found below.

Backend Microservices
---
api - Serves JSON data to be ingested by React JS frontend application. To build
    the project run `docker compose up --build api`. The accessible ports can
    be edited in `docker-compose.yml`.

datapull - Pulls data from Chicago Data Portal using Go and the SODA API. To
    build the project run `docker-compose build --no-cache datapull` or omit the
    `--no-cache` flag to speed things up a bit. To run the service use
    `docker-compose up datapull`.

postgres - Database for the data lake. To run the database use
    `docker-compose up postgres`.

Scripts
---
Scripts for development are stored in the "scripts" directory and can generally
be run with bash. For example, `bash scripts/rebuild-containers.sh`.
