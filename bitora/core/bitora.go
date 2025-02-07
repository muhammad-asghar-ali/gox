package core

import (
	"github.com/boltdb/bolt"
)

type (
	Bitora struct {
		Tip []byte
		DB  *bolt.DB
	}

	BitoraIterator struct {
		Current []byte
		DB      *bolt.DB
	}
)

const dbFile = "bitora.db"
const bucket = "blocks"

func (bit *Bitora) AddBlock(data string) {
	var last []byte

	err := bit.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		last = b.Get([]byte("l"))

		return nil
	})
	if err != nil {
		return
	}

	new := NewBlock(data, last)
	err = bit.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		if err := b.Put(new.Hash, new.Serialize()); err != nil {
			return nil
		}

		err = b.Put([]byte("l"), new.Hash)
		bit.Tip = new.Hash

		return nil
	})
}

// the first in the chain, is called genesis block
func NewGenesisBlock() *Block {
	return NewBlock("Genesis Block", []byte{})
}

func NewBitora() *Bitora {
	var tip []byte
	db, _ := bolt.Open(dbFile, 0600, nil)

	err := db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))

		if b == nil {
			genesis := NewGenesisBlock()
			b, err := tx.CreateBucket([]byte(bucket))
			if err != nil {
				return nil
			}

			b.Put(genesis.Hash, genesis.Serialize())
			b.Put([]byte("l"), genesis.Hash)
			tip = genesis.Hash
		} else {
			tip = b.Get([]byte("l"))
		}

		return nil
	})
	if err != nil {
		return nil
	}

	return &Bitora{tip, db}
}

func (bit *Bitora) Iterator() *BitoraIterator {
	bci := &BitoraIterator{bit.Tip, bit.DB}

	return bci
}

func (i *BitoraIterator) Next() *Block {
	block := &Block{}

	err := i.DB.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucket))
		encodedBlock := b.Get(i.Current)
		block = Deserialize(encodedBlock)

		return nil
	})

	if err != nil {
		return nil
	}

	i.Current = block.PrevHash

	return block
}
