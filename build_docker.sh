#!/bin/bash

HELLO_VERSION="1.0"

docker build -t devokun/hello . 

if [ "x${PUSH_TO_REGISTRY}" == "x" ]; then
  echo "Not pushing to Docker Registry"
  exit 0
fi

echo "Logging in to Registry as $DOCKER_REGISTRY_USERNAME"
docker login \
  -u ${DOCKER_REGISTRY_USERNAME} \
  -p ${DOCKER_REGISTRY_PASSWORD}

IMAGE_ID=$( docker inspect devokun/hello --format '{{.Id}}' | sed 's/^.*://g' )

echo "Tagging Image $IMAGE_ID"
docker tag ${IMAGE_ID} devokun/hello:latest
docker tag ${IMAGE_ID} devokun/hello:${HELLO_VERSION}

echo "Pushing Image $IMAGE_ID to Registry"
docker push devokun/hello:latest
docker push devokun/hello:${HELLO_VERSION}
