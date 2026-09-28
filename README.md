# Echo Service

[![CI](https://github.com/rodrigodlopes/echo-service-takehome/actions/workflows/ci.yml/badge.svg)](https://github.com/rodrigodlopes/echo-service-takehome/actions/workflows/ci.yml)

A Go web service that echoes every request back as JSON (headers, query params, body, path), packaged as a Docker image and deployed to a local [Kind](https://kind.sigs.k8s.io/) cluster with Pulumi (Go, local state).

```
$ curl -X POST "http://localhost:8888/hello?x=1" -H "X-Test: value" -d 'hi'
{"Headers":{"Accept":["*/*"],"X-Test":["value"],...},"Params":{"x":["1"]},"Body":"hi","Path":"/hello"}
```

## Layout

- `echo-service/`: the Go app, its unit test, and the `Dockerfile`
- `echo-infra/`: the Pulumi program (a `Deployment` and a ClusterIP `Service`)
- `scripts/ci.sh`: runs the tests, builds the image, and loads it into Kind
- `.github/workflows/ci.yml`: runs `ci.sh`, deploys with Pulumi, and smoke-tests the endpoint on every push

## Getting started

See [SETUP.md](SETUP.md).
