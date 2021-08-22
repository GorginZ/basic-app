#!/usr/bin/env bash
# ecrname=$(aws cloudformation describe-stacks --stack-name basic-app-ecr --query "Stacks[0].Outputs[0].OutputValue" --output text)

# commit=$(git describe --tags --always)
# docker build -t "$ecrname:$commit" . 
# docker push "$ecrname:$commit"
# docker build -t "$ecrname:latest" . 
# docker push "$ecrname:latest"


docker build --build-arg AWS_ACCESS_KEY_ID="${{ secrets.AWS_ACCESS_KEY_ID }}" \
    --build-arg AWS_SECRET_ACCESS_KEY="${{ secrets.AWS_SECRET_ACCESS_KEY }}" \
    --build-arg AWS_DEFAULT_REGION="${{ secrets.AWS_DEFAULT_REGION }}" \
    -t $ECR_REGISTRY/$ECR_REPOSITORY:$IMAGE_TAG .
docker push $ECR_REGISTRY/$ECR_REPOSITORY:$IMAGE_TAG
echo "::set-output name=image::$ECR_REGISTRY/$ECR_REPOSITORY:$IMAGE_TAG"
