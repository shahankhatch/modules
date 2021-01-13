module github.com/okwme/modules/incubator/faucet

go 1.15

replace github.com/cosmos/cosmos-sdk => /Users/shahank/git_interchain/cosmos-sdk
replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4

require (
	github.com/cosmos/cosmos-sdk v0.40.0
	github.com/gorilla/mux v1.8.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.1
	github.com/tmdvs/Go-Emoji-Utils v1.1.0
)

//replace github.com/cosmos/cosmos-sdk v0.38.4 => github.com/okwme/cosmos-sdk v0.38.5-0.20200715162801-4fd244eef297

// replace github.com/cosmos/cosmos-sdk v0.38.4 => /Users/billy/GitHub.com/okwme/cosmos-sdk
