package crypton

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"os"
)

func ToHax(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	HandleError(err)

	return buff.Bytes()
}

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func CoinbaseTx(to, data string) *Transection {
	if data == "" {
		data = fmt.Sprintf("Coins to %s", to)
	}

	txin := &TxInput{[]byte{}, -1, data}
	txout := &TxOutput{100, to}

	tx := &Transection{nil, []*TxInput{txin}, []*TxOutput{txout}}
	tx.SetID()

	return tx
}

func DatabaseExists() bool {
	if _, err := os.Stat(db_file); os.IsNotExist(err) {
		return false
	}

	return true
}
