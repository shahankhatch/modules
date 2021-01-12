package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
)

// ModuleCdc is the codec for the module
var ModuleCdc *codec.AminoCodec
var Cdc *codec.LegacyAmino

func init() {
	RegisterCodec(ModuleCdc)
}

// RegisterCodec registers concrete types on the Amino codec
func RegisterCodec(cdc *codec.AminoCodec) {
	cdc.RegisterConcrete(MsgMint{}, "faucet/Mint", nil)
	cdc.RegisterConcrete(MsgFaucetKey{}, "faucet/FaucetKey", nil)
}
