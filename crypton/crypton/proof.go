package crypton

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const difficulty = 18

type (
	ProofOfWork struct {
		Block  *Block
		Target *big.Int
	}
)

func (pow *ProofOfWork) InitData(nonce int) []byte {
	return bytes.Join(
		[][]byte{
			pow.Block.PrevHash,
			pow.Block.HashTransections(),
			ToHax(int64(nonce)),
			ToHax(int64(difficulty)),
		},
		[]byte{},
	)
}

func (pow *ProofOfWork) Run() (int, []byte) {
	var intHash big.Int
	hash := [32]byte{}
	nonce := 0

	for nonce < math.MaxInt64 {
		data := pow.InitData(nonce)
		hash = sha256.Sum256(data)

		fmt.Printf("\r %x", hash)
		intHash.SetBytes(hash[:])

		if intHash.Cmp(pow.Target) == -1 {
			break
		} else {
			nonce++
		}
	}

	fmt.Println()
	return nonce, hash[:]
}

func (pow *ProofOfWork) Validate() bool {
	var intHash big.Int

	data := pow.InitData(pow.Block.Nonce)
	hash := sha256.Sum256(data)

	return intHash.SetBytes(hash[:]).Cmp(pow.Target) == -1
}

func NewProof(b *Block) *ProofOfWork {
	t := big.NewInt(1)
	t.Lsh(t, uint(256-difficulty))

	return &ProofOfWork{b, t}
}
