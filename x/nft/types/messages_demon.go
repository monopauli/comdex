package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
	"unicode/utf8"
)

const (
	TypeMsgCreateDenom   = "create_denom"
	TypeMsgUpdateDenom   = "update_denom"
	TypeMsgTransferDenom = "transfer_denom"
	TypeMsgMintNFT       = "mint_nft"
	TypeMsgEditNFT       = "edit_nft"
	TypeMsgTransferNFT   = "transfer_nft"
	TypeMsgBurnNFT       = "burn_nft"
)

var (
	_ sdk.Msg = &MsgCreateDenom{}
	_ sdk.Msg = &MsgUpdateDenom{}
	_ sdk.Msg = &MsgTransferDenom{}
	_ sdk.Msg = &MsgMintNFT{}
	_ sdk.Msg = &MsgEditNFT{}
	_ sdk.Msg = &MsgTransferNFT{}
	_ sdk.Msg = &MsgBurnNFT{}
)

func NewMsgCreateDenom(symbol, name, schema, description, previewUri, sender string) *MsgCreateDenom {
	return &MsgCreateDenom{
		Sender:      sender,
		Id:          GenUniqueID(DenomPrefix),
		Symbol:      symbol,
		Name:        name,
		Schema:      schema,
		Description: description,
		PreviewURI:  previewUri,
	}
}

func (msg MsgCreateDenom) Route() string { return RouterKey }

func (msg MsgCreateDenom) Type() string { return TypeMsgCreateDenom }

func (msg MsgCreateDenom) ValidateBasic() error {
	if err := ValidateDenomID(msg.Id); err != nil {
		return err
	}
	if err := ValidateDenomSymbol(msg.Symbol); err != nil {
		return err
	}
	name := strings.TrimSpace(msg.Name)
	if len(name) > 0 && !utf8.ValidString(name) {
		return sdkerrors.Wrap(ErrInvalidDenom, "denom name is invalid")
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCreateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgCreateDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgUpdateDenom(id, name, description, previewUri, sender string) *MsgUpdateDenom {
	return &MsgUpdateDenom{
		Id:          id,
		Name:        name,
		Description: description,
		PreviewURI:  previewUri,
		Sender:      sender,
	}
}

func (msg MsgUpdateDenom) Route() string { return RouterKey }

func (msg MsgUpdateDenom) Type() string { return TypeMsgUpdateDenom }

func (msg MsgUpdateDenom) ValidateBasic() error {
	if err := ValidateDenomID(msg.Id); err != nil {
		return err
	}
	name := msg.Name
	if len(name) > 0 && !utf8.ValidString(name) {
		return sdkerrors.Wrap(ErrInvalidDenom, "denom name is invalid")
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgUpdateDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgUpdateDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgTransferDenom(id, sender, recipient string) *MsgTransferDenom {
	return &MsgTransferDenom{
		Id:        id,
		Sender:    sender,
		Recipient: recipient,
	}
}

func (msg MsgTransferDenom) Route() string { return RouterKey }

func (msg MsgTransferDenom) Type() string { return TypeMsgTransferDenom }

func (msg MsgTransferDenom) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address; %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address; %s", err)
	}
	return nil
}

func (msg MsgTransferDenom) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgTransferDenom) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgMintNFT(
	denomId, sender, recipient string, metadata Metadata, data string, transferable, extensible bool) *MsgMintNFT {

	return &MsgMintNFT{
		Id:           GenUniqueID(IDPrefix),
		DenomId:      denomId,
		Metadata:     metadata,
		Data:         data,
		Transferable: transferable,
		Extensible:   extensible,
		Sender:       sender,
		Recipient:    recipient,
	}
}

func (msg MsgMintNFT) Route() string { return RouterKey }

func (msg MsgMintNFT) Type() string { return TypeMsgMintNFT }

func (msg MsgMintNFT) ValidateBasic() error {

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address; %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address; %s", err)
	}

	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if err := ValidateMediaURI(msg.Metadata.MediaURI); err != nil {
		return err
	}

	return ValidateNFTID(msg.Id)
}

func (msg MsgMintNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgMintNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgTransferNFT(id, denomId, sender, recipient string) *MsgTransferNFT {

	return &MsgTransferNFT{
		Id:        strings.ToLower(strings.TrimSpace(id)),
		DenomId:   strings.TrimSpace(denomId),
		Sender:    sender,
		Recipient: recipient,
	}
}

func (msg MsgTransferNFT) Route() string { return RouterKey }

func (msg MsgTransferNFT) Type() string { return TypeMsgTransferNFT }

func (msg MsgTransferNFT) ValidateBasic() error {
	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address; %s", err)
	}

	if _, err := sdk.AccAddressFromBech32(msg.Recipient); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid recipient address; %s", err)
	}
	return ValidateNFTID(msg.Id)
}

func (msg MsgTransferNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgEditNFT(
	id, denomId string, metadata Metadata, data,
	transferable, extensible, sender string) *MsgEditNFT {
	return &MsgEditNFT{
		Id:           id,
		DenomId:      denomId,
		Metadata:     metadata,
		Data:         data,
		Transferable: transferable,
		Extensible:   extensible,
		Sender:       sender,
	}
}

func (msg MsgEditNFT) Route() string { return RouterKey }

func (msg MsgEditNFT) Type() string { return TypeMsgEditNFT }

func (msg MsgEditNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address; %s", err)
	}

	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}

	if err := ValidateMediaURI(msg.Metadata.MediaURI); err != nil {
		return err
	}
	return ValidateNFTID(msg.Id)
}

func (msg MsgEditNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgEditNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}

func NewMsgBurnNFT(denomId, id, sender string) *MsgBurnNFT {
	return &MsgBurnNFT{
		DenomId: denomId,
		Id:      id,
		Sender:  sender,
	}
}

func (msg MsgBurnNFT) Route() string { return RouterKey }

func (msg MsgBurnNFT) Type() string { return TypeMsgBurnNFT }

func (msg MsgBurnNFT) ValidateBasic() error {
	if _, err := sdk.AccAddressFromBech32(msg.Sender); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address; %s", err)
	}

	if err := ValidateDenomID(msg.DenomId); err != nil {
		return err
	}
	return ValidateNFTID(msg.Id)
}

func (msg MsgBurnNFT) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgBurnNFT) GetSigners() []sdk.AccAddress {
	from, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{from}
}
