package core

type (
	Bitora struct {
		Blocks []*Block
	}
)

func (bit *Bitora) AddBlock(data string) {
	prev := bit.Blocks[len(bit.Blocks)-1]
	new := NewBlock(data, prev.Hash)

	bit.Blocks = append(bit.Blocks, new)
}

// the first in the chain, is called genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBitora() *Bitora {
	return &Bitora{[]*Block{NewGenesisBlock()}}
}
