#!/bin/bash
ACTION=${action}
if [ "$ACTION" == "run" ];
then
  ENVIRONMENT=develop go run ./cmd/gocicd/index.go
elif [ "$ACTION" == "test" ];
then
  go test -cover ./cmd/gocicd/domain ./cmd/gocicd/infra ./cmd/gocicd/interface/controllers
elif [ "$ACTION" == "unit-test" ];
then
  go test -cover ./cmd/gocicd/domain
elif [ "$ACTION" == "integration-test" ];
then
  go test -cover ./cmd/gocicd/interface/controllers
elif [ "$ACTION" == "e2e-test" ];
then
  if [ "$TRAVIS_BRANCH" == "" ]; then
    CURRENT_HOST=http://localhost:$PORT
  else
    CURRENT_HOST=$(gcloud run services describe go-cicd-$TRAVIS_BRANCH-$TRAVIS_BUILD_ID  --platform managed --region us-east1 --format 'value(status.url)')
  fi
  HOST=$CURRENT_HOST go test ./tests/end2end
fi