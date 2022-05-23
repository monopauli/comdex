package cli

import (
	"github.com/comdex-official/comdex/x/liquidation/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
	"strconv"
)

func txWhitelistAppId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "whitelist-app-id [app_mapping_Id]",
		Short: "Add Whitelisted appId for liquidations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			appMappingId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgWhitelistAppId(
				appMappingId,
				ctx.GetFromAddress(),
			)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}

func txRemoveWhitelistAppId() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "remove-whitelist-app-id [app_mapping_Id] ",
		Short: "Remove Whitelisted appId for liquidations",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			appMappingId, err := strconv.ParseUint(args[0], 10, 64)
			if err != nil {
				return err
			}

			msg := types.NewMsgRemoveWhitelistAsset(
				appMappingId,
				ctx.GetFromAddress(),
			)

			return tx.GenerateOrBroadcastTxCLI(ctx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd

}