package auth

import (
	"encoding/json"
	sdk "github.com/cosmos/amino-js/go/lib/cosmos/cosmos-sdk/types"
)

var _ sdk.Tx = (*StdTx)(nil)

type StdTx struct {
	Msgs       []sdk.Msg      `json:"msg"`
	Fee        StdFee         `json:"fee"`
	Signatures []StdSignature `json:"signatures"`
	Memo       string         `json:"memo"`
}

type StdFee struct {
	Amount sdk.Coins `json:"amount"`
	Gas    uint64    `json:"gas"`
}

type StdSignDoc struct {
	AccountNumber uint64            `json:"account_number"`
	ChainID       string            `json:"chain_id"`
	Fee           json.RawMessage   `json:"fee"`
	Memo          string            `json:"memo"`
	Msgs          []json.RawMessage `json:"msgs"`
	Sequence      uint64            `json:"sequence"`
}

type StdSignature struct {
	PubKey    []byte `json:"pub_key"`
	Signature []byte `json:"signature"`
}
