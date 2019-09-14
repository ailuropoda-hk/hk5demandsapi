#!/bin/sh

DOCKER_FILE=./docker/development/Dockerfile
PROJECT_NAME=hk5demandsapi-dev
LAST_COMMIT=$(git rev-parse --short HEAD 2> /dev/null | sed "s/\(.*\)/\1/")
IMAGE_TAG=${PROJECT_NAME}:${LAST_COMMIT}

echo "Docker Build Push Details"
echo "Image Tag:   ${IMAGE_TAG}"
echo "Docker File: ${DOCKER_FILE}"

read -p "Are you sure? " -n 1 -r
echo
if [[ $REPLY =~ ^[Yy]$ ]]
then
  echo "Deploying..."
  cd ../..
  docker build --rm -f ${DOCKER_FILE} -t ${IMAGE_TAG} .
else
  echo "Deployment Cancelled"
fi