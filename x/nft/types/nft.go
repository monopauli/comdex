package types

import (
	"github.com/comdex-official/comdex/x/nft/exported"
	"strings"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ exported.NFT = NFT{}

func NewNFT(
	id string, metadata Metadata, data string, transferable, extensible bool, owner sdk.AccAddress,
	createdTime time.Time) NFT {
	return NFT{
		Id:           id,
		Metadata:     metadata,
		Data:         data,
		Owner:        owner.String(),
		Transferable: transferable,
		Extensible:   extensible,
		CreatedAt:    createdTime,
	}
}

func (nft NFT) GetID() string {
	return nft.Id
}

func (nft NFT) GetName() string {
	return nft.Metadata.Name
}

func (nft NFT) GetDescription() string {
	return nft.Metadata.Description
}

func (nft NFT) GetMediaURI() string {
	return nft.Metadata.MediaURI
}

func (nft NFT) GetPreviewURI() string {
	return nft.Metadata.PreviewURI
}

func (nft NFT) GetOwner() sdk.AccAddress {
	owner, _ := sdk.AccAddressFromBech32(nft.Owner)
	return owner
}

func (nft NFT) GetMetadata() string {
	return nft.Metadata.String()
}
func (nft NFT) GetData() string {
	return nft.Data
}
func (nft NFT) IsTransferable() bool {
	return nft.Transferable
}
func (nft NFT) IsExtensible() bool {
	return nft.Extensible
}
func (nft NFT) GetCreatedTime() time.Time {
	return nft.CreatedAt
}

// NFT

type NFTs []exported.NFT

func NewNFTs(nfts ...exported.NFT) NFTs {
	if len(nfts) == 0 {
		return NFTs{}
	}
	return nfts
}

func ValidateNFTID(nftID string) error {
	nftID = strings.TrimSpace(nftID)
	if len(nftID) < MinDenomLen || len(nftID) > MaxDenomLen {
		return sdkerrors.Wrapf(
			ErrInvalidNFTID,
			"invalid nftID %s, only accepts value [%d, %d]", nftID, MinDenomLen, MaxDenomLen)
	}
	if !IsBeginWithAlpha(nftID) || !IsAlphaNumeric(nftID) {
		return sdkerrors.Wrapf(
			ErrInvalidNFTID,
			"invalid nftID %s, only accepts alphanumeric characters,and begin with an english letter", nftID)
	}
	return nil
}

func ValidateMediaURI(mediaURI string) error {
	if len(mediaURI) > MaxMediaURILen {
		return sdkerrors.Wrapf(ErrInvalidMediaURI, "invalid mediaURI %s, only accepts value [0, %d]", mediaURI, MaxMediaURILen)
	}
	return nil
}

func ValidatePreviewURI(previewURI string) error {
	if len(previewURI) > MaxPreviewURILen {
		return sdkerrors.Wrapf(ErrInvalidPreviewURI, "invalid previewURI %s, only accepts value [0, %d]", previewURI, MaxPreviewURILen)
	}
	return nil
}
