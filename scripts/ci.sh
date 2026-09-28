#!/usr/bin/env bash
set -euo pipefail

IMAGE_NAME="echo-service:latest"
KIND_CLUSTER="${KIND_CLUSTER:-echo-cluster}"

cd "$(dirname "$0")/../echo-service"

echo "==> Running unit tests"
go test ./...

echo "==> Building Docker image ($IMAGE_NAME)"
docker build -t "$IMAGE_NAME" .

echo "==> Loading image into kind cluster ($KIND_CLUSTER)"
kind load docker-image "$IMAGE_NAME" --name "$KIND_CLUSTER"

echo "==> Done"
