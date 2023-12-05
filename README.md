![build](https://github.com/orltom/sandbox/actions/workflows/ci.yaml/badge.svg)
![tilt](https://github.com/orltom/sandbox/actions/workflows/tilt.yaml/badge.svg)

# Intention
This project is a personal playground to learn about new technologies and concepts. 
It intends to explore all facets of the application lifecycle, from development to operations.

# Project Idea
The software project idea is to develop a platform that randomly publishes jokes.

## Prerequisites
* Install [golang](https://go.dev/doc/install)
* Install [npm](https://nodejs.org/en/download/package-manager)
* Install [docker](https://www.docker.com/)
* Install [kubectl](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/)
* Install [k3d](https://k3d.io/)
* Install [tilt](https://tilt.dev/)

## Development Setup
```shell
# create k8s cluster
k3d cluster create --config config/k3d/dev-cluster.yaml

# Build and deploy
tilt up
```

## Contributing
Please use the GitHub issue tracker to submit bugs or request features.

## Disclaimer
Copyright Orlando Tomás.

Distributed under the terms of the MIT license, tool is free and open source software.