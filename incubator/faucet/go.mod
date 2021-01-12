module github.com/okwme/modules/incubator/faucet

go 1.13

require (
	github.com/bartekn/go-bip39 v0.0.0-20171116152956-a05967ea095d // indirect
	github.com/cosmos/cosmos-sdk v0.40.0-rc5
	github.com/gorilla/mux v1.8.0
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/iavl v0.13.2 // indirect
	github.com/tendermint/tendermint v0.34.0
	github.com/tmdvs/Go-Emoji-Utils v1.1.0
)

replace github.com/cosmos/cosmos-sdk v0.38.4 => github.com/okwme/cosmos-sdk v0.38.5-0.20200715162801-4fd244eef297

// replace github.com/cosmos/cosmos-sdk v0.38.4 => /Users/billy/GitHub.com/okwme/cosmos-sdk
