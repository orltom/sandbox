load('ext://helm_resource', 'helm_resource', 'helm_repo')
load('ext://cert_manager', 'deploy_cert_manager')

analytics_settings(enable=False)

allow_k8s_contexts('default')

include('./tilt/demo.tilt')


