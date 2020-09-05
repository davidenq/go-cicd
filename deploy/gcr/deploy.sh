#!/bin/bash

gcloud run deploy go-cicd-${TRAVIS_BRANCH} --image ${CONTAINER_TAG}:${TRAVIS_COMMIT} --region us-east1 --platform managed --allow-unauthenticated