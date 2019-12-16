#!/bin/sh

IMAGE_NAME="$1"
kustomize edit set image app-image="${IMAGE_NAME}"
kustomize build . > ../../manifest/prd-k8s.yaml
