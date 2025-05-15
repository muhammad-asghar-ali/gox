package crypton

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
)

type (
	Block struct {
		Hash         []byte
		Transections []*Transection
		PrevHash     []byte
		Nonce        int
	}
)

func (b *Block) Create(txs []*Transection, prev []byte) *Block {
	block := &Block{[]byte{}, txs, prev, 0}

	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce
	return block
}

func (b *Block) Genesis(coinbase *Transection) *Block {
	return b.Create([]*Transection{coinbase}, []byte{})
}

func (b *Block) Serialize() []byte {
	res := bytes.Buffer{}
	encoder := gob.NewEncoder(&res)

	err := encoder.Encode(b)
	HandleError(err)

	return res.Bytes()
}

func (b *Block) Deserialize(data []byte) *Block {
	block := &Block{}

	decoder := gob.NewDecoder(bytes.NewReader(data))
	err := decoder.Decode(block)
	HandleError(err)

	return block
}

func (b *Block) HashTransections() []byte {
	txsh := [][]byte{}
	hash := [32]byte{}

	for _, tx := range b.Transections {
		txsh = append(txsh, tx.ID)
	}

	hash = sha256.Sum256(bytes.Join(txsh, []byte{}))
	return hash[:]
}
