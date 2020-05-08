package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

type MsgTransfer struct {
	SourcePort    string         `json:"source_port" yaml:"source_port"`       // the port on which the packet will be sent
	SourceChannel string         `json:"source_channel" yaml:"source_channel"` // the channel by which the packet will be sent
	DestHeight    uint64         `json:"dest_height" yaml:"dest_height"`       // the current height of the destination chain
	Amount        sdk.Coins      `json:"amount" yaml:"amount"`                 // the tokens to be transferred
	Sender        sdk.AccAddress `json:"sender" yaml:"sender"`                 // the sender address
	Receiver      string         `json:"receiver" yaml:"receiver"`             // the recipient address on the destination chain
}

type FungibleTokenPacketData struct {
	Amount   sdk.Coins `json:"amount" yaml:"amount"`     // the tokens to be transferred
	Sender   string    `json:"sender" yaml:"sender"`     // the sender address
	Receiver string    `json:"receiver" yaml:"receiver"` // the recipient address on the destination chain
}