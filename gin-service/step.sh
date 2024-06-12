#!/bin/bash
git pull
docker stop go-container
docker rm go-container
docker rmi go-image
go build -o main-dev
docker build . -t go-image
docker run -dit --name go-container --privileged=true -p 13001:13001 go-image