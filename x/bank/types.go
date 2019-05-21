package bank

import (
	"bytes"
	cosmos "github.com/cosmos/cosmos-sdk/types"
	"github.com/regen-network/regen-ledger/x/action"
	abci "github.com/tendermint/tendermint/types"
)

type SendCapability struct {
	Account    cosmos.AccAddress
	SpendLimit cosmos.Coins
	//DailySpendLimit cosmos.Coins
	//SpentToday      cosmos.Coins
}

func (cap SendCapability) CapabilityKey() string {
	return "bank.send"
}

func (cap SendCapability) RootAccount() cosmos.AccAddress {
	return cap.Account
}

func (cap SendCapability) Accept(msg action.Action, block abci.Header) (allow bool, updated action.Capability, delete bool) {
	switch msg := msg.(type) {
	case MsgSend:
		if bytes.Equal(msg.From, cap.Account) {
			left, valid := cap.SpendLimit.SafeMinus(msg.Amount)
			if !valid {
				return false, nil, false
			}
			if left.IsZero() {
				return true, nil, true
			}
			return true, SendCapability{Account: cap.Account, SpendLimit: left}, false
		}
	}
	return false, nil, false
}

type MsgSend struct {
	From   cosmos.AccAddress
	To     cosmos.AccAddress
	Amount cosmos.Coins
}

func (msg MsgSend) Route() string {
	panic("implement me")
}

func (msg MsgSend) Type() string {
	panic("implement me")
}

func (msg MsgSend) ValidateBasic() cosmos.Error {
	panic("implement me")
}

func (msg MsgSend) GetSignBytes() []byte {
	panic("implement me")
}

func (msg MsgSend) GetSigners() []cosmos.AccAddress {
	panic("implement me")
}

func (msg MsgSend) RequiredCapabilities() []action.Capability {
	panic("implement me")
}
