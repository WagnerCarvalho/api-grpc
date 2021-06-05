#!/bin/bash

cp ../../authentication/authsvc .
cp ../../api/apisvc .

docker build -t qagile/api-grpc:v1 .
docker inspect qagile/api-grpc:v1