#!/bin/sh
cd `dirname $0`

IMAGE_NAME="$1"
kustomize edit set image client-deployment="${IMAGE_NAME}"
kustomize build . > ../../../manifest/production/client-deploy.yaml
