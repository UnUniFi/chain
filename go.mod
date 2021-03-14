module github.com/lcnem/jpyx

go 1.16

require (
	github.com/KimuraYu45z/test v0.0.0-20210225023031-f3624bc07db8 // indirect
	github.com/cosmos/cosmos-sdk v0.42.1
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/spf13/cast v1.3.1
	github.com/spf13/cobra v1.1.3
	github.com/spf13/pflag v1.0.5
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.8
	github.com/tendermint/tm-db v0.6.4
	go.etcd.io/etcd v0.0.0-20191023171146-3cf2f69b5738
	google.golang.org/genproto v0.0.0-20210312152112-fc591d9ea70f
	google.golang.org/grpc v1.36.0
	gopkg.in/yaml.v2 v2.4.0 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

// patch bech32 decoding to enable larger string lengths
replace github.com/btcsuite/btcutil => github.com/kava-labs/btcutil v0.0.0-20200522184203-886d33430f06
