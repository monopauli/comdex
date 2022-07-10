package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func NewMsgDeposit(depositor string, appID uint64, amount sdk.Coin) *MsgDepositESM {
	return &MsgDepositESM{
		Depositor: depositor,
		AppId:     appID,
		Amount:    amount,
	}
}

func (msg MsgDepositESM) Route() string { return ModuleName }

func (msg *MsgDepositESM) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetDepositor())
	if err != nil {
		return err
	}

	if asset := msg.GetAmount(); !asset.IsValid() {
		return sdkerrors.Wrap(ErrInvalidAsset, asset.String())
	}

	return nil
}

func (msg *MsgDepositESM) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetDepositor())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on.
func (msg *MsgDepositESM) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func NewMsgExecute(depositor string, appID uint64) *MsgExecuteESM {
	return &MsgExecuteESM{
		Depositor: depositor,
		AppId:     appID,
	}
}

func (msg MsgExecuteESM) Route() string { return ModuleName }

func (msg *MsgExecuteESM) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.GetDepositor())
	if err != nil {
		return err
	}

	return nil
}

func (msg *MsgExecuteESM) GetSigners() []sdk.AccAddress {
	lender, _ := sdk.AccAddressFromBech32(msg.GetDepositor())
	return []sdk.AccAddress{lender}
}

// GetSignBytes get the bytes for the message signer to sign on.
func (msg *MsgExecuteESM) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}