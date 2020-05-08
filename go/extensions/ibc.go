package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

// MsgTransfer defines a msg to transfer fungible tokens (i.e Coins) between ICS20 enabled chains.
// See ICS Spec here: https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#data-structures
type MsgTransfer struct {
	SourcePort    string         `json:"source_port" yaml:"source_port"`       // the port on which the packet will be sent
	SourceChannel string         `json:"source_channel" yaml:"source_channel"` // the channel by which the packet will be sent
	DestHeight    uint64         `json:"dest_height" yaml:"dest_height"`       // the current height of the destination chain
	Amount        sdk.Coins      `json:"amount" yaml:"amount"`                 // the tokens to be transferred
	Sender        sdk.AccAddress `json:"sender" yaml:"sender"`                 // the sender address
	Receiver      string         `json:"receiver" yaml:"receiver"`             // the recipient address on the destination chain
}

// FungibleTokenPacketData defines a struct for the packet payload
// See FungibleTokenPacketData spec: https://github.com/cosmos/ics/tree/master/spec/ics-020-fungible-token-transfer#data-structures
type FungibleTokenPacketData struct {
	Amount   sdk.Coins `json:"amount" yaml:"amount"`     // the tokens to be transferred
	Sender   string    `json:"sender" yaml:"sender"`     // the sender address
	Receiver string    `json:"receiver" yaml:"receiver"` // the recipient address on the destination chain
}
