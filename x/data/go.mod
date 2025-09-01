module github.com/regen-network/regen-ledger/x/data/v3

go 1.21

require (
	cosmossdk.io/api v0.9.2
	cosmossdk.io/store v1.1.1
	cosmossdk.io/errors v1.0.1
	github.com/cosmos/cosmos-sdk v0.50.14
	github.com/cometbft/cometbft v0.38.15
	github.com/cometbft/cometbft-db v0.14.1
	github.com/cosmos/cosmos-db v1.1.1
	github.com/cosmos/btcutil v1.0.5
	github.com/cosmos/gogoproto v1.7.0
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.4
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/regen-network/gocuke v1.1.0
	github.com/regen-network/regen-ledger/api/v2 v2.4.0
	github.com/regen-network/regen-ledger/types/v2 v2.0.0
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.9.0
	golang.org/x/crypto v0.21.0
	google.golang.org/genproto/googleapis/api v0.0.0-20240123012728-ef4313101c80
	google.golang.org/grpc v1.74.2
	google.golang.org/protobuf v1.36.7
	gotest.tools/v3 v3.5.1
)


replace (
	github.com/99designs/keyring => github.com/cosmos/keyring v1.2.0
	// dgrijalva/jwt-go is deprecated and doesn't receive security updates.
	github.com/dgrijalva/jwt-go => github.com/golang-jwt/jwt/v4 v4.4.2
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	github.com/syndtr/goleveldb => github.com/syndtr/goleveldb v1.0.1-0.20210819022825-2ae1ddf74ef7
	// stick with compatible version or x/exp in v0.47.x line for must be v0.0.0-20230711153332-06a737ee72cb for gogoproto v1.4.10
	golang.org/x/exp => golang.org/x/exp v0.0.0-20230711153332-06a737ee72cb
)

replace (
	github.com/regen-network/regen-ledger/api/v2 => ../../api
	github.com/regen-network/regen-ledger/types/v2 => ../../types
	github.com/regen-network/regen-ledger/orm => ../../orm
)
