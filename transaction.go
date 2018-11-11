package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"fmt"
	"log"
)

const subsidy = 10

// Transaction is include UTXO and UTXI
type Transaction struct {
	ID   []byte
	Vin  []TXInput
	Vout []TXOutput
}

// SetID sets Id of a transaction
func (tx *Transaction) SetID() {
	var encoded bytes.Buffer
	var hash [32]byte

	enc := gob.NewEncoder(&encoded)
	err := enc.Encode(tx)
	if err != nil {
		log.Panic(err)
	}
	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

// TXOutput save gogocoin(Value) and ScriptPubKey is
type TXOutput struct {
	Value        int
	scriptPubKey string
}

// TXInput refers to previous TXOutput
type TXInput struct {
	Txid      []byte
	Vout      int
	ScriptSig string
}

// NewCoinbaseTX is create Transaction
func NewCoinbaseTX(to, data string) *Transaction {
	if data == "" {
		data = fmt.Sprintf("Reward to '%s'", to)
	}
	txin := TXInput{[]byte{}, -1, data}
	txout := TXOutput{subsidy, to}
	tx := Transaction{nil, []TXInput{txin}, []TXOutput{txout}}
	tx.SetID()
	return &tx
}
