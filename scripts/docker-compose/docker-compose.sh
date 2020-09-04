#!/bin/bash

docker-compose -p gocicd --env-file=./cmd/gocicd/.env --project-directory=./ -f ./deploy/docker-compose/docker-compose.yaml ${a}