#!/usr/bin/env bash
set -e
ecrname=$(aws cloudformation describe-stacks --stack-name basic-app-ecr --query "Stacks[0].Outputs[0].OutputValue" --output text)

commit=$(git describe --tags --always)
docker build -t "$ecrname:$commit" . 
docker push "$ecrname:$commit"
docker build -t "$ecrname:latest" . 
docker push "$ecrname:latest"

