General Architecture
---
This project is divided into several microservices that are packaged with Docker
. Instructions for building and running the individual microservices can be
found below.

Backend Microservices
---
datapull - Pulls data from Chicago Data Portal using Go and the SODA API. To
    build the project run `docker-compose build --no-cache datapull` or omit the
    `--no-cache` flag to speed things up a bit. To run the service use
    `docker-compose up datapull`
Geocode - Geocoding and reverse geocoding with Go and Google Geocoder
Postgres - Database for the data lake
ReportGenerator - Builds reports from

Building and Deploying
---
