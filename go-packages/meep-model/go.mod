module github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-model

go 1.12

require (
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis v0.0.0
	github.com/RyanCarrier/dijkstra v0.0.0-20190726134004-b51cadb5ae52
	github.com/RyanCarrier/dijkstra-1 v0.0.0-20170512020943-0e5801a26345 // indirect
	github.com/albertorestifo/dijkstra v0.0.0-20160910063646-aba76f725f72 // indirect
	github.com/blang/semver v3.5.1+incompatible
	github.com/google/uuid v1.2.0
	github.com/mattomatic/dijkstra v0.0.0-20130617153013-6f6d134eb237 // indirect
)

replace (
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr => ../../go-packages/meep-data-key-mgr
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model => ../../go-packages/meep-data-model
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger => ../../go-packages/meep-logger
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis => ../../go-packages/meep-redis
)
