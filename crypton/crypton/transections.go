package crypton

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"encoding/hex"
	"log"
)

type (
	TxInput struct {
		ID  []byte
		Out int
		Sig string
	}

	TxOutput struct {
		Value  int
		PubKey string
	}

	Transection struct {
		ID      []byte
		Inputs  []*TxInput
		Outputs []*TxOutput
	}
)

func (tx *Transection) SetID() {
	encoded := bytes.Buffer{}
	hash := [32]byte{}

	encode := gob.NewEncoder(&encoded)
	err := encode.Encode(tx)
	HandleError(err)

	hash = sha256.Sum256(encoded.Bytes())
	tx.ID = hash[:]
}

func (tx *Transection) IsCoinbase() bool {
	return len(tx.Inputs) == 1 && len(tx.Inputs[0].ID) == 0 && tx.Inputs[0].Out == -1
}

func (in *TxInput) CanUnlock(data string) bool {
	return in.Sig == data
}

func (out *TxOutput) CanBeUnlock(data string) bool {
	return out.PubKey == data
}

func NewTransection(from, to string, amount int, c *Crypton) *Transection {
	inputs := []*TxInput{}
	outputs := []*TxOutput{}

	acc, valid_outs := c.FindSpendableOutputs(from, amount)
	if acc < amount {
		log.Panic("not enough funds")
	}

	for txid, outs := range valid_outs {
		txID, err := hex.DecodeString(txid)
		HandleError(err)

		for _, out := range outs {
			input := &TxInput{txID, out, from}
			inputs = append(inputs, input)
		}

	}

	outputs = append(outputs, &TxOutput{amount, to})
	if acc > amount {
		outputs = append(outputs, &TxOutput{acc - amount, from})
	}

	tx := &Transection{nil, inputs, outputs}
	tx.SetID()

	return tx
}
