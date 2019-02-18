module gitlab.com/regen-network/regen-ledger

require (
	github.com/DATA-DOG/godog v0.7.10
	github.com/ZondaX/hid-go v0.4.0
	github.com/ZondaX/ledger-go v0.4.0
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/beorn7/perks v0.0.0-20180321164747-3a771d992973
	github.com/bgentry/speakeasy v0.1.0
	github.com/btcsuite/btcd v0.0.0-20181130015935-7d2daa5bfef2
	github.com/btcsuite/btcutil v0.0.0-20180706230648-ab6388e0c60a
	github.com/cosmos/cosmos-sdk v0.32.0
	github.com/cosmos/go-bip39 v0.0.0-20180618194314-52158e4697b8
	github.com/davecgh/go-spew v1.1.1
	github.com/fsnotify/fsnotify v1.4.7
	github.com/go-kit/kit v0.6.0
	github.com/go-logfmt/logfmt v0.4.0
	github.com/go-stack/stack v1.8.0
	github.com/gogo/protobuf v1.1.1
	github.com/golang/protobuf v1.2.0
	github.com/golang/snappy v0.0.0-20180518054509-2e65f85255db
	github.com/gorilla/context v1.1.1
	github.com/gorilla/mux v1.6.2
	github.com/gorilla/websocket v1.4.0
	github.com/hashicorp/hcl v1.0.0
	github.com/inconshreveable/mousetrap v1.0.0
	github.com/jmhodges/levigo v0.0.0-20161115193449-c42d9e0ca023
	github.com/kr/logfmt v0.0.0-20140226030751-b84e30acd515
	github.com/magiconair/properties v1.8.0
	github.com/mattn/go-isatty v0.0.4
	github.com/matttproud/golang_protobuf_extensions v1.0.1
	github.com/mitchellh/mapstructure v1.1.2
	github.com/pelletier/go-toml v1.2.0
	github.com/pkg/errors v0.8.0
	github.com/pmezard/go-difflib v1.0.0
	github.com/prometheus/client_golang v0.9.1
	github.com/prometheus/client_model v0.0.0-20180712105110-5c3871d89910
	github.com/prometheus/common v0.0.0-20181126121408-4724e9255275
	github.com/prometheus/procfs v0.0.0-20181129180645-aa55a523dc0a
	github.com/rakyll/statik v0.1.4
	github.com/rcrowley/go-metrics v0.0.0-20180503174638-e2704e165165
	github.com/rs/cors v1.6.0
	github.com/spf13/afero v1.1.2
	github.com/spf13/cast v1.3.0
	github.com/spf13/cobra v0.0.3
	github.com/spf13/jwalterweatherman v1.0.0
	github.com/spf13/pflag v1.0.3
	github.com/spf13/viper v1.0.3
	github.com/stretchr/testify v1.2.2
	github.com/syndtr/goleveldb v0.0.0-20181128100959-b001fa50d6b2
	github.com/tendermint/btcd v0.0.0-20180816174608-e5840949ff4f
	github.com/tendermint/go-amino v0.14.1
	github.com/tendermint/iavl v0.12.0
	github.com/tendermint/tendermint v0.30.0
	github.com/twpayne/go-geom v1.0.4
	github.com/zondax/ledger-cosmos-go v0.9.2
	golang.org/x/crypto v0.0.0-20190211182817-74369b46fc67
	golang.org/x/net v0.0.0-20190213061140-3a22650c66bd
	golang.org/x/sys v0.0.0-20181128092732-4ed8d59d0b35
	golang.org/x/text v0.3.1-0.20180807135948-17ff2d5776d2
	google.golang.org/genproto v0.0.0-20190201180003-4b09977fb922
	google.golang.org/grpc v1.17.0
	gopkg.in/yaml.v2 v2.2.2
)

replace golang.org/x/crypto => github.com/tendermint/crypto v0.0.0-20180820045704-3764759f34a5

// replace github.com/DATA-DOG/godog => github.com/regen-network/godog v0.0.0-0.20190215155814-31cb7bc0e9a6bdc4c3116e54eecbc6ce453d9a2b
