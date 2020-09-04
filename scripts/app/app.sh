#!/bin/bash
ACTION=${a}
if [ "$ACTION" == "run" ];
then
  ENVIRONMENT=develop go run ./cmd/gocicd/index.go
elif [ "$ACTION" == "test" ];
then
  go test -cover ./...
elif [ "$ACTION" == "unit-test" ];
then
  go test -cover ./cmd/gocicd/domain
elif [ "$ACTION" == "integration-test" ];
then
  go test -cover ./cmd/gocicd/interface/controllers
elif [ "$ACTION" == "e2e-test" ];
then
  echo "e2e.test"
fi