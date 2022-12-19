load('ext://cert_manager', 'deploy_cert_manager')

analytics_settings(enable=False)

allow_k8s_contexts('default')

include('./tilt/database.tilt')
include('./tilt/api.tilt')

include('./tilt/telemetry.tilt')

