# Simple Web Application Demo
A simple demo web application written in Golang. 

## Prerequisites
* Install [docker](https://www.docker.com/)
* Install [golang](https://go.dev/doc/install)
* Install [kubectl](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/)
* Install [k3d](https://k3d.io/)
* Install [tilt](https://tilt.dev/)

## Setup
Setup k3d cluster
```shell
k3d cluster create --config k3d/dev-cluster.yaml
```

## Usage
Build and deploy
```shell
tilt up
```

Test REST endpoint
```shell
curl localhost:8081/api/v1/jokes/random
```

## Contributing
Please use the GitHub issue tracker to submit bugs or request features.

## Disclaimer
Copyright Orlando Tomás.

Distributed under the terms of the MIT license, tool is free and open source software.