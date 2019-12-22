#!/bin/sh
cd `dirname $0`

IMAGE_NAME="$1"
kustomize edit set image server-deployment="${IMAGE_NAME}"
kustomize build . > ../../../manifest/production/server-deploy.yaml
