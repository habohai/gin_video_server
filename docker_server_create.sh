#! /bin/bash

# create the docker server images
cd ./apiserver
docker build -t apiserver .
cd ..

cd ./scheduler
docker build -t scheduler .
cd ..

cd ./streamserver
docker build -t streamserver .
cd ..

cd ./web
docker build -t web .
cd ..