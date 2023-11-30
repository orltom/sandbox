# Idea
This project is a personal playground to learn about new technologies and concepts. 
It intends to explore all facets of the application lifecycle, from development to operations.

## Prerequisites
* Install [golang](https://go.dev/doc/install)
* Install [npm](https://nodejs.org/en/download/package-manager)
* Install [docker](https://www.docker.com/)
* Install [kubectl](https://kubernetes.io/de/docs/tasks/tools/install-kubectl/)
* Install [k3d](https://k3d.io/)
* Install [tilt](https://tilt.dev/)

## Development Setup
Setup k3d cluster
```shell
k3d cluster create --config config/k3d/dev-cluster.yaml
```

Build and deploy
```shell
tilt up
```

Open browser 
```shell
http://127.0.0.1.nip.io/ui
```

## Contributing
Please use the GitHub issue tracker to submit bugs or request features.

## Disclaimer
Copyright Orlando Tomás.

Distributed under the terms of the MIT license, tool is free and open source software.