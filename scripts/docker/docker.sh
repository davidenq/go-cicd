#!/bin/bash
ACTION=${a}
if [ "$ACTION" == "run" ];
then
  docker run -e ENVIRONMENT=${ENVIRONMENT}  -e PORT=${PORT} -p ${PORT}:${PORT} gocicd:latest
elif [ "$ACTION" == "build" ];
then
  docker build  --file=./build/package/docker/Dockerfile --compress --force-rm --rm --tag  gocicd .
fi