package core

import (
	"bytes"
	"crypto/sha256"
	"encoding/gob"
	"strconv"
	"time"
)

type (
	Block struct {
		Timestamp int64
		Data      []byte
		PrevHash  []byte
		Hash      []byte
		Nonce     int
	}
)

func (b *Block) SetHash() {
	timestamp := []byte(strconv.FormatInt(b.Timestamp, 10))
	headers := bytes.Join([][]byte{b.PrevHash, b.Data, timestamp}, []byte{})
	hash := sha256.Sum256(headers)

	b.Hash = hash[:]
}

func (b *Block) Serialize() []byte {
	res := &bytes.Buffer{}
	encoder := gob.NewEncoder(res)

	if err := encoder.Encode(b); err != nil {
		return nil
	}

	return res.Bytes()
}

func Deserialize(d []byte) *Block {
	b := &Block{}

	decoder := gob.NewDecoder(bytes.NewReader(d))
	if err := decoder.Decode(b); err != nil {
		return nil
	}

	return b
}

func NewBlock(data string, prevBlockHash []byte) *Block {
	block := &Block{time.Now().Unix(), []byte(data), prevBlockHash, []byte{}, 0}
	pow := NewProofOfWork(block)
	nonce, hash := pow.Run()

	block.Hash = hash[:]
	block.Nonce = nonce

	return block
}
