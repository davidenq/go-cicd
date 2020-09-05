#!/bin/bash

gcloud run deploy go-cicd-${TRAVIS_BRANCH}-${TRAVIS_BUILD_ID} --image ${CONTAINER_TAG}:${TRAVIS_COMMIT} --region us-east1 --platform managed --allow-unauthenticated