# Setup

## Prerequisites

- Go 1.27+
- Docker
- [kind](https://kind.sigs.k8s.io/)
- [Pulumi CLI](https://www.pulumi.com/docs/install/)
- kubectl

## 1. Create the local Kind cluster

```
kind create cluster --name echo-cluster
```

## 2. Build, test, and load the image into the cluster

```
./scripts/ci.sh
```

This runs the Go unit tests, builds the `echo-service:latest` Docker image, and loads it into the `echo-cluster` Kind cluster (`kind load docker-image`), so it's available without a registry.

## 3. Deploy with Pulumi

Pulumi is configured to use a local filesystem state backend (no cloud account needed).

```
cd echo-infra
pulumi login --local
export PULUMI_CONFIG_PASSPHRASE=""   # local dev only, no real secrets stored
pulumi stack init dev                # first time only
pulumi up
```

This creates a `Deployment` and a `Service` (`echo-service`, ClusterIP, port 80 -> 8080) in the cluster.

## 4. Try it

```
kubectl port-forward svc/echo-service 8888:80
curl -X POST "http://localhost:8888/hello?x=1" -H "X-Test: value" -d '{"k":"v"}'
```

Returns a JSON body with `Headers`, `Params`, `Body`, and `Path`.

## Teardown

```
cd echo-infra && pulumi destroy
kind delete cluster --name echo-cluster
```
