package cli

import (
	"fmt"

	// "strings"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	sdk "github.com/cosmos/cosmos-sdk/types"

	// "github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	// sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/okwme/modules/incubator/faucet/internal/keeper"
	"github.com/okwme/modules/incubator/faucet/internal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.AminoCodec) *cobra.Command {
	// Group pooltoy queries under a subcommand
	pooltoyQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	brrrCommand := GetCmdWhenBrrr(queryRoute, cdc)
	flags.AddTxFlagsToCmd(brrrCommand)
	pooltoyQueryCmd.AddCommand(
		brrrCommand,
	)

	return pooltoyQueryCmd
}

func GetCmdWhenBrrr(queryRoute string, cdc *codec.AminoCodec) *cobra.Command {
	return &cobra.Command{
		Use:   "when-brrr [userAccount]",
		Short: "how many seconds until this user can brrr again",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			address := args[0]
			_, err := sdk.AccAddressFromBech32(address)
			if err != nil {
				fmt.Printf("could not query User\n%s\n", err.Error())
				return nil
			}
			res, _, err := clientCtx.QueryWithData(fmt.Sprintf("custom/%s/%s/%s", queryRoute, keeper.QueryWhenBrrr, address), nil)
			if err != nil {
				fmt.Printf("could not query User\n%s\n", err.Error())
				return nil
			}
			var out int64
			clientCtx.LegacyAmino.MustUnmarshalJSON(res, &out)
			return clientCtx.PrintObjectLegacy(out)
		},
	}
}
