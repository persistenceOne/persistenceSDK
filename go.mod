module github.com/persistenceOne/persistenceSDK

go 1.14

require (
	github.com/99designs/keyring v1.1.6
	github.com/CosmWasm/wasmd v0.15.1
	github.com/Shopify/sarama v1.19.0
	github.com/asaskevich/govalidator v0.0.0-20200108200545-475eaeb16496
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d
	github.com/cosmos/cosmos-sdk v0.42.6
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/gogo/protobuf v1.3.3
	github.com/gorilla/mux v1.8.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/onsi/ginkgo v1.14.2 // indirect
	github.com/onsi/gomega v1.10.4 // indirect
	github.com/pkg/errors v0.9.1
	github.com/regen-network/cosmos-proto v0.3.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15
	github.com/tendermint/tendermint v0.34.11
	github.com/tendermint/tm-db v0.6.4
	golang.org/x/text v0.3.4 // indirect
	golang.org/x/xerrors v0.0.0-20200804184101-5ec99f83aff1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	honnef.co/go/tools v0.0.1-2020.1.6
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
