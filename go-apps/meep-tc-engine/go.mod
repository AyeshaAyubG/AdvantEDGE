module github.com/InterDigitalInc/AdvantEDGE/go-apps/meep-tc-engine

go 1.12

require (
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mg-manager-model v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-model v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mq v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-net-char-mgr v0.0.0
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis v0.0.0
	github.com/gogo/protobuf v1.2.1 // indirect
	github.com/google/btree v1.0.0 // indirect
	github.com/google/gofuzz v1.0.0 // indirect
	github.com/googleapis/gnostic v0.2.0 // indirect
	github.com/gregjones/httpcache v0.0.0-20190212212710-3befbb6ad0cc // indirect
	github.com/json-iterator/go v1.1.6 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.1 // indirect
	github.com/peterbourgon/diskv v2.0.1+incompatible // indirect
	github.com/spf13/pflag v1.0.3 // indirect
	golang.org/x/oauth2 v0.0.0-20190402181905-9f3314589c9a // indirect
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4 // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/api v0.0.0-20181204000039-89a74a8d264d // indirect
	k8s.io/apimachinery v0.0.0-20181127025237-2b1284ed4c93
	k8s.io/client-go v10.0.0+incompatible
	k8s.io/klog v0.0.0-20181108234604-8139d8cb77af // indirect
	sigs.k8s.io/yaml v1.1.0 // indirect
)

replace (
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-key-mgr => ../../go-packages/meep-data-key-mgr
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-data-model => ../../go-packages/meep-data-model
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-gis-cache => ../../go-packages/meep-gis-cache
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-logger => ../../go-packages/meep-logger
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mg-manager-model => ../../go-packages/meep-mg-manager-model
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-model => ../../go-packages/meep-model
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-mq => ../../go-packages/meep-mq
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-net-char-mgr => ../../go-packages/meep-net-char-mgr
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-pdu-session-store => ../../go-packages/meep-pdu-session-store
	github.com/InterDigitalInc/AdvantEDGE/go-packages/meep-redis => ../../go-packages/meep-redis
)
