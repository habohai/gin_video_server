#! /bin/bash

# clean the docker images

docker-compose stop
docker-compose down
docker rmi -f scheduler apiserver web streamserver
#docker container prune
#docker image prune

