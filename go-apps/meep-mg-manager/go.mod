module github.com/InterDigitalInc/AdvantEDGE/go-apps/meep-mg-manager

go 1.12

require (
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-http-logger v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mg-app-client v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mg-manager-model v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-model v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mq v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-sessions v0.0.0
	github.com/RyanCarrier/dijkstra v0.0.0-20190726134004-b51cadb5ae52
	github.com/gorilla/handlers v1.4.0
	github.com/gorilla/mux v1.7.4
)

replace (
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr => ../../go-packages/meep-data-key-mgr
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model => ../../go-packages/meep-data-model
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-http-logger => ../../go-packages/meep-http-logger
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger => ../../go-packages/meep-logger
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-metric-store => ../../go-packages/meep-metric-store
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mg-app-client => ../../go-packages/meep-mg-app-client
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mg-manager-model => ../../go-packages/meep-mg-manager-model
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-model => ../../go-packages/meep-model
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mq => ../../go-packages/meep-mq
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis => ../../go-packages/meep-redis
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-sessions => ../../go-packages/meep-sessions
)
