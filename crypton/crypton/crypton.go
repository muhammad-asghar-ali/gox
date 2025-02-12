package crypton

import (
	"encoding/hex"
	"fmt"
	"runtime"

	"github.com/dgraph-io/badger/v4"
)

const (
	path    = "./tmp/blocks"
	db_file = "./tmp/blocks/MANIFEST"
	genesis = "First Transection from Genesis"
)

type (
	Crypton struct {
		LastHash []byte
		Database *badger.DB
	}

	CryptonIterator struct {
		CurrentHash []byte
		Database    *badger.DB
	}
)

func (c *Crypton) AddBlock(txs []*Transection) {
	b := &Block{}
	lh := []byte{}

	err := c.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		HandleError(err)

		err = item.Value(func(val []byte) error {
			lh = append([]byte{}, val...)
			return nil
		})

		return err
	})

	HandleError(err)

	new := b.Create(txs, lh)
	err = c.Database.Update(func(txn *badger.Txn) error {
		err := txn.Set(new.Hash, new.Serialize())
		HandleError(err)

		err = txn.Set([]byte("lh"), new.Hash)
		c.LastHash = new.Hash

		return err
	})

	HandleError(err)
}

func (c *Crypton) Init(addr string) *Crypton {
	if DatabaseExists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	b := &Block{}
	lh := []byte{}

	opts := badger.DefaultOptions(path)

	db, err := badger.Open(opts)
	HandleError(err)

	err = db.Update(func(txn *badger.Txn) error {
		fmt.Println("No existing blockchain found")
		genesis := b.Genesis(CoinbaseTx(addr, genesis))

		fmt.Println("Genesis Proved")
		err = txn.Set(genesis.Hash, genesis.Serialize())
		HandleError(err)

		err = txn.Set([]byte("lh"), genesis.Hash)

		lh = genesis.Hash

		return err
	})

	HandleError(err)

	return &Crypton{lh, db}
}

func (c *Crypton) Contiune(adds string) *Crypton {
	if DatabaseExists() {
		fmt.Println("Blockchain already exists")
		runtime.Goexit()
	}

	lh := []byte{}

	opts := badger.DefaultOptions(path)

	db, err := badger.Open(opts)
	HandleError(err)

	err = db.Update(func(txn *badger.Txn) error {
		item, err := txn.Get([]byte("lh"))
		HandleError(err)

		err = item.Value(func(val []byte) error {
			lh = append([]byte{}, val...)
			return nil
		})

		return err
	})

	HandleError(err)

	chain := &Crypton{lh, db}
	return chain
}

func (c *Crypton) Iterator() *CryptonIterator {
	return &CryptonIterator{c.LastHash, c.Database}
}

func (iter *CryptonIterator) Next() *Block {
	block := &Block{}
	err := iter.Database.View(func(txn *badger.Txn) error {
		item, err := txn.Get(iter.CurrentHash)
		HandleError(err)

		encoded := []byte{}
		err = item.Value(func(val []byte) error {
			encoded = append([]byte{}, val...)
			return nil
		})

		b := &Block{}
		block = b.Deserialize(encoded)
		return err
	})

	HandleError(err)

	iter.CurrentHash = block.PrevHash

	return block
}

func (c *Crypton) FindUnspentTxs(addr string) []*Transection {
	utxs := make([]*Transection, 0) // unspent

	spent_txos := make(map[string][]int) // txos - tx outputs

	iter := c.Iterator()

	for {
		b := iter.Next()
		for _, tx := range b.Transections {
			txID := hex.EncodeToString(tx.ID)

		Outputs:
			for oidx, out := range tx.Outputs {
				if spent_txos[txID] != nil {
					for _, sptout := range spent_txos[txID] {
						if sptout == oidx {
							continue Outputs
						}
					}
				}

				if out.CanBeUnlock(addr) {
					utxs = append(utxs, tx)
				}
			}

			if tx.IsCoinbase() == false {
				for _, in := range tx.Inputs {
					if in.CanUnlock(addr) {
						in_txID := hex.EncodeToString(in.ID)
						spent_txos[in_txID] = append(spent_txos[in_txID], in.Out)
					}
				}
			}
		}

		if len(b.PrevHash) == 0 {
			break
		}
	}

	return utxs
}

func (c *Crypton) FindUTXO(addr string) []*TxOutput {
	utxos := make([]*TxOutput, 0)
	unspt_txs := c.FindUnspentTxs(addr)

	for _, tx := range unspt_txs {
		for _, out := range tx.Outputs {
			if out.CanBeUnlock(addr) {
				utxos = append(utxos, out)
			}
		}
	}

	return utxos
}

func (c *Crypton) FindSpendableOutputs(addr string, amount int) (int, map[string][]int) {
	unspt_outs := make(map[string][]int)
	unspt_txs := c.FindUnspentTxs(addr)

	acc := 0

Work:
	for _, tx := range unspt_txs {
		txID := hex.EncodeToString(tx.ID)

		for oidx, out := range tx.Outputs {
			if out.CanBeUnlock(addr) && acc < amount {
				acc += out.Value
				unspt_outs[txID] = append(unspt_outs[txID], oidx)

				if acc >= amount {
					break Work
				}
			}
		}
	}

	return acc, unspt_outs
}
