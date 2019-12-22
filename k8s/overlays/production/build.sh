#!/bin/sh
cd `dirname $0`

CLIENT_IMAGE_NAME="$1"
SERVER_IMAGE_NAME="$2"
kustomize edit set image client-deployment="${CLIENT_IMAGE_NAME}"
kustomize edit set image server-deployment="${SERVER_IMAGE_NAME}"
kustomize build . > ../../manifest/production/deploy.yaml
