package cli

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/version"
	"strings"

	"github.com/comdex-official/comdex/x/nft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"
)

// NewTxCmd returns the transaction commands for this module
func NewTxCmd() *cobra.Command {
	txCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "NFT transactions subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	txCmd.AddCommand(
		GetCmdCreateDenom(),
		GetCmdUpdateDenom(),
		GetCmdTransferDenom(),
		GetCmdMintNFT(),
		GetCmdEditNFT(),
		GetCmdTransferNFT(),
		GetCmdBurnNFT(),
	)

	return txCmd
}

func GetCmdCreateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use: "create [symbol]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Create a new denom.
Example:
$ %s tx nft create [symbol] --name=<name> --schema=<schema> --description=<description> --preview-uri=<preview-uri> 
--chain-id=<chain-id> --from=<key-name> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			symbol := args[0]

			denomName, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}

			schema, err := cmd.Flags().GetString(FlagSchema)
			if err != nil {
				return err
			}

			description, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			previewURI, err := cmd.Flags().GetString(FlagPreviewURI)
			if err != nil {
				return err
			}

			msg := types.NewMsgCreateDenom(symbol,
				denomName,
				schema,
				description,
				previewURI,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsCreateDenom)
	_ = cmd.MarkFlagRequired(FlagName)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdMintNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use: "mint [denom-id]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Mint an NFT.
Example:
$ %s tx nft mint [denom-id] --type <nft-type> --name <nft-name> --description <nft-descritpion> --media-uri=<uri> --preview-uri=<uri> 
--transferable yes --extensible yes --recipient=<recipient> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			denomId := args[0]

			sender := clientCtx.GetFromAddress().String()

			recipient, err := cmd.Flags().GetString(FlagRecipient)
			if err != nil {
				return err
			}

			if len(recipient) > 0 {
				if _, err = sdk.AccAddressFromBech32(recipient); err != nil {
					return err
				}
			} else {
				recipient = sender
			}

			nftMetadata := types.Metadata{}
			nftName, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}

			nftDescription, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			nftMediaURI, err := cmd.Flags().GetString(FlagMediaURI)
			if err != nil {
				return err
			}

			nftPreviewURI, err := cmd.Flags().GetString(FlagPreviewURI)
			if err != nil {
				return err
			}

			if len(nftName) > 0 {
				nftMetadata.Name = nftName
			}
			if len(nftDescription) > 0 {
				nftMetadata.Description = nftDescription
			}
			if len(nftMediaURI) > 0 {
				nftMetadata.MediaURI = nftMediaURI
			}
			if len(nftPreviewURI) > 0 {
				nftMetadata.PreviewURI = nftPreviewURI
			}
			data, err := cmd.Flags().GetString(FlagData)
			if err != nil {
				return err
			}

			var transferable bool
			transferability, err := cmd.Flags().GetString(FlagTransferable)
			if err != nil {
				return err
			}
			transferability = strings.ToLower(transferability)
			if transferability == "false" || transferability == "no" {
				transferable = false
			} else if transferability == "true" || transferability == "yes" {
				transferable = true
			} else {
				return fmt.Errorf("invalid option for transferable flag , valid options are true|false, yes|no")
			}
			var extensible bool
			extensibility, err := cmd.Flags().GetString(FlagExtensible)
			if err != nil {
				return err
			}
			if extensibility == "false" || extensibility == "no" {
				extensible = false
			} else if extensibility == "true" || extensibility == "yes" {
				extensible = true
			} else {
				return fmt.Errorf("invalid option for extensible flag , valid options are yes|no")
			}

			msg := types.NewMsgMintNFT(
				denomId,
				sender,
				recipient,
				nftMetadata,
				data,
				transferable,
				extensible,
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsMintNFT)
	_ = cmd.MarkFlagRequired(FlagMediaURI)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdEditNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use: "edit [denom-id] [nft-id]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Edit the data of an NFT.
Example:
$ %s tx nft edit [denom-id] [nft-id] --name=<nft-name> --description=<nft-description> --media-uri=<uri>
--preview-uri=<uri> --type=<nft-type> --transferable=yes --extensible=yes --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			denomId := strings.ToLower(strings.TrimSpace(args[0]))
			nftId := strings.ToLower(strings.TrimSpace(args[1]))

			nftMetadata := types.Metadata{}
			nftName, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}
			nftName = strings.TrimSpace(nftName)

			nftDescription, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}
			nftDescription = strings.TrimSpace(nftDescription)

			nftMediaURI, err := cmd.Flags().GetString(FlagMediaURI)
			if err != nil {
				return err
			}
			nftMediaURI = strings.TrimSpace(nftMediaURI)

			nftPreviewURI, err := cmd.Flags().GetString(FlagPreviewURI)
			if err != nil {
				return err
			}
			nftPreviewURI = strings.TrimSpace(nftPreviewURI)

			if len(nftName) > 0 {
				nftMetadata.Name = nftName
			}
			if len(nftDescription) > 0 {
				nftMetadata.Description = nftDescription
			}
			if len(nftMediaURI) > 0 {
				nftMetadata.MediaURI = nftMediaURI
			}
			if len(nftPreviewURI) > 0 {
				nftMetadata.PreviewURI = nftPreviewURI
			}
			data, err := cmd.Flags().GetString(FlagData)
			if err != nil {
				return err
			}

			transferable, err := cmd.Flags().GetString(FlagTransferable)
			if err != nil {
				return err
			}
			if !(len(transferable) > 0 && (transferable == "no" || transferable == "yes" ||
				transferable == types.DoNotModify)) {
				return fmt.Errorf("invalid option for transferable flag , valid options are yes | no")
			}
			extensible, err := cmd.Flags().GetString(FlagExtensible)
			if err != nil {
				return err
			}
			if !(len(extensible) > 0 && (extensible == "no" || extensible == "yes" || extensible == types.DoNotModify)) {
				return fmt.Errorf("invalid option for extensible flag , valid options are yes|no")
			}
			msg := types.NewMsgEditNFT(
				nftId,
				denomId,
				nftMetadata,
				data,
				transferable,
				extensible,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsEditNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdUpdateDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use: "update-denom [denom-id]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Edit the data of Denom.
Example:
$ %s tx nft update-denom [denom-id] --name=<nft-name> --description=<nft-description> 
--preview-uri=<uri> --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			denomId := strings.TrimSpace(args[0])

			denomName, err := cmd.Flags().GetString(FlagName)
			if err != nil {
				return err
			}

			denomDescription, err := cmd.Flags().GetString(FlagDescription)
			if err != nil {
				return err
			}

			denomPreviewURI, err := cmd.Flags().GetString(FlagPreviewURI)
			if err != nil {
				return err
			}

			msg := types.NewMsgUpdateDenom(
				denomId,
				denomName,
				denomDescription,
				denomPreviewURI,
				clientCtx.GetFromAddress().String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsUpdateDenom)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdTransferDenom() *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfer-denom [recipient] [denom-id]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Transfer a denom to a recipient.
Example:
$ %s tx nft transfer-denom [recipient] [denom-id] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			denomId := args[1]

			msg := types.NewMsgTransferDenom(
				denomId,
				clientCtx.GetFromAddress().String(),
				recipient.String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsTransferDenom)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdTransferNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use: "transfer [recipient] [denom-id] [nft-id]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Transfer an NFT to a recipient.
Example:
$ %s tx nft transfer [recipient] [denom-id] [nft-id] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			recipient, err := sdk.AccAddressFromBech32(args[0])
			if err != nil {
				return err
			}

			denomId := strings.ToLower(strings.TrimSpace(args[1]))
			nftId := strings.ToLower(strings.TrimSpace(args[2]))

			msg := types.NewMsgTransferNFT(
				nftId,
				denomId,
				clientCtx.GetFromAddress().String(),
				recipient.String(),
			)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	cmd.Flags().AddFlagSet(FsTransferNFT)
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}

func GetCmdBurnNFT() *cobra.Command {
	cmd := &cobra.Command{
		Use: "burn [denom-id] [nft-id]",
		Long: strings.TrimSpace(
			fmt.Sprintf(`Burn an NFT.
Example:
$ %s tx nft burn [denom-id] [nft-id] --from=<key-name> --chain-id=<chain-id> --fees=<fee>`,
				version.AppName,
			),
		),
		Args: cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			denomId := strings.ToLower(strings.TrimSpace(args[0]))
			nftId := strings.ToLower(strings.TrimSpace(args[1]))

			msg := types.NewMsgBurnNFT(denomId, nftId, clientCtx.GetFromAddress().String())
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)

	return cmd
}
