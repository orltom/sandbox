load('ext://helm_resource', 'helm_resource', 'helm_repo')
load('ext://cert_manager', 'deploy_cert_manager')

analytics_settings(enable=False)

allow_k8s_contexts('default')

local_resource(
  dir='.',
  name='build',
  cmd='make build-api',
  deps=['cmd'],
  labels=["demo"],
  allow_parallel=False
)

docker_build(
  ref="golang-http-example:latest",
  context=".",
  dockerfile="config/docker/httpd/Dockerfile",
  build_args={'TARGETARCH': 'amd64'},
  only=[
    'bin/golang-http-example-amd64',
  ],
)

k8s_yaml([
    'config/k8s/httpd/deployment.yaml',
    'config/k8s/httpd/service.yaml',
])

k8s_resource(
  workload='api',
  labels=["demo"]
)

docker_build(
  ref="golang-db-example:latest",
  context=".",
  dockerfile="config/docker/database/Dockerfile",
)

k8s_yaml([
    'config/k8s/database/deployment.yaml',
    'config/k8s/database/service.yaml',
])

k8s_resource(
  workload='database',
  labels=["demo"]
)


