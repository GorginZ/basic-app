#!/usr/bin/env bash

name='basic-app'
aws cloudformation deploy \
      --template-file ./cfn/ecr.yaml \
      --parameter-overrides \
            RepositoryName=$name \
      --stack-name $name-ecr \
      --region ap-southeast-2 \
      --no-fail-on-empty-changeset
