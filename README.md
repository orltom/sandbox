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
k3d cluster create dev-cluster \
  --agents-memory=8G \
  --registry-create dev-cluster-registry \
  -p "30000-30099:30000-30099@server:0"
```

## Usage
```shell
tilt up
```

## Contributing
Please use the GitHub issue tracker to submit bugs or request features.

## Disclaimer
Copyright Orlando Tomás.

Distributed under the terms of the MIT license, tool is free and open source software.