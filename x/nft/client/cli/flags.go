package cli

import (
	flag "github.com/spf13/pflag"
)

const (
	FlagName         = "name"
	FlagDescription  = "description"
	FlagMediaURI     = "media-uri"
	FlagPreviewURI   = "preview-uri"
	FlagData         = "data"
	FlagTransferable = "transferable"
	FlagExtensible   = "extensible"
	FlagRecipient    = "recipient"
	FlagOwner        = "owner"
	FlagDenomID      = "denom-id"
	FlagSchema       = "schema"
)

var (
	FsCreateDenom   = flag.NewFlagSet("", flag.ContinueOnError)
	FsUpdateDenom   = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferDenom = flag.NewFlagSet("", flag.ContinueOnError)
	FsMintNFT       = flag.NewFlagSet("", flag.ContinueOnError)
	FsEditNFT       = flag.NewFlagSet("", flag.ContinueOnError)
	FsTransferNFT   = flag.NewFlagSet("", flag.ContinueOnError)
	FsQuerySupply   = flag.NewFlagSet("", flag.ContinueOnError)
	FsQueryOwner    = flag.NewFlagSet("", flag.ContinueOnError)
)

func init() {
	FsCreateDenom.String(FlagSchema, "", "Denom schema")
	FsCreateDenom.String(FlagName, "", "Name of the denom")
	FsCreateDenom.String(FlagDescription, "", "Description for denom")
	FsCreateDenom.String(FlagPreviewURI, "", "Preview image uri for denom")

	FsUpdateDenom.String(FlagName, "[do-not-modify]", "Name of the denom")
	FsUpdateDenom.String(FlagDescription, "[do-not-modify]", "Description for denom")
	FsUpdateDenom.String(FlagPreviewURI, "[do-not-modify]", "Preview image uri for denom")

	FsTransferDenom.String(FlagRecipient, "", "recipient of the denom")

	FsMintNFT.String(FlagMediaURI, "", "Media uri of nft")
	FsMintNFT.String(FlagRecipient, "", "Receiver of the nft. default value is sender address of transaction")
	FsMintNFT.String(FlagPreviewURI, "", "Preview uri of nft")
	FsMintNFT.String(FlagName, "", "Name of nft")
	FsMintNFT.String(FlagDescription, "", "Description of nft")
	FsMintNFT.String(FlagData, "", "custom data of nft")

	FsMintNFT.String(FlagTransferable, "yes", "transferability of nft (yes | no)")
	FsMintNFT.String(FlagExtensible, "yes", "extensisbility of nft (yes | no)")

	FsEditNFT.String(FlagMediaURI, "[do-not-modify]", "Media uri of nft")
	FsEditNFT.String(FlagPreviewURI, "[do-not-modify]", "Preview uri of nft")
	FsEditNFT.String(FlagName, "[do-not-modify]", "Name of nft")
	FsEditNFT.String(FlagDescription, "[do-not-modify]", "Description of nft")
	FsEditNFT.String(FlagTransferable, "[do-not-modify]", "transferability of nft")
	FsEditNFT.String(FlagData, "[do-not-modify]", "custom data of nft")
	FsEditNFT.String(FlagExtensible, "[do-not-modify]", "extensibility of nft (yes | no)")

	FsTransferNFT.String(FlagRecipient, "", "Receiver of the nft. default value is sender address of transaction")
	FsQuerySupply.String(FlagOwner, "", "The owner of a nft")
	FsQueryOwner.String(FlagDenomID, "", "id of the denom")
}
