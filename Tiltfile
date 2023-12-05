update_settings(k8s_upsert_timeout_secs=120)

ci_settings(timeout = '30m')

analytics_settings(enable=False)

load('ext://cert_manager', 'deploy_cert_manager')

allow_k8s_contexts('default')

include('./infrastructure/tilt/database.tilt')
include('./infrastructure/tilt/api.tilt')
include('./infrastructure/tilt/ui.tilt')

include('./infrastructure/tilt/telemetry.tilt')

