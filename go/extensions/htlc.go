package extensions

import (
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

type MsgCreateHTLC struct {
	Sender               sdk.AccAddress `json:"sender"`                  // the initiator address
	To                   sdk.AccAddress `json:"to"`                      // the destination address
	ReceiverOnOtherChain string         `json:"receiver_on_other_chain"` // the claim receiving address on the other chain
	Amount               sdk.Coins      `json:"amount"`                  // the amount to be transferred
	HashLock             []byte         `json:"hash_lock"`               // the hash lock generated from secret (and timestamp if provided)
	Timestamp            uint64         `json:"timestamp"`               // if provided, used to generate the hash lock together with secret
	TimeLock             uint64         `json:"time_lock"`               // the time span after which the HTLC will expire
}

type MsgClaimHTLC struct {
	Sender   sdk.AccAddress `json:"sender"`    // the initiator address
	HashLock []byte         `json:"hash_lock"` // the hash lock identifying the HTLC to be claimed
	Secret   []byte         `json:"secret"`    // the secret with which to claim
}

type MsgRefundHTLC struct {
	Sender   sdk.AccAddress `json:"sender"`    // the initiator address
	HashLock []byte         `json:"hash_lock"` // the hash lock identifying the HTLC to be refunded
}