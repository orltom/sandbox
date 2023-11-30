load('ext://cert_manager', 'deploy_cert_manager')

analytics_settings(enable=False)

allow_k8s_contexts('default')

include('./infrastructure/tilt/database.tilt')
include('./infrastructure/tilt/api.tilt')
include('./infrastructure/tilt/ui.tilt')

include('./infrastructure/tilt/telemetry.tilt')

