#!/bin/sh

PROJECT_NAME=hk5demandsapi-dev
LAST_COMMIT=$(git rev-parse --short HEAD 2> /dev/null | sed "s/\(.*\)/\1/")
IMAGE_TAG=${PROJECT_NAME}:${LAST_COMMIT}


cd ../..

docker container rm $PROJECT_NAME
docker run --name $PROJECT_NAME -v $(pwd):/go/src/hk5demandsapi -v $1:/go/src/hk5demandsapi/data/visualdata -p 8080:8080 -it $IMAGE_TAG  /bin/bash