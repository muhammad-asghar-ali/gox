package tigerbeetle

import (
	"sync"

	. "github.com/tigerbeetle/tigerbeetle-go"
	. "github.com/tigerbeetle/tigerbeetle-go/pkg/types"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/helpers"
)

type (
	Tb struct {
		client Client
	}
)

var (
	instance *Tb
	once     sync.Once
)

func Set() *Tb {
	once.Do(func() {
		client, err := NewClient(ToUint128(0), []string{"3000"})
		helpers.HandleError(err)

		instance = &Tb{client: client}
	})

	return instance
}

func GetTb() *Tb {
	if instance == nil {
		return Set()
	}

	return instance
}

func (c *Tb) Close() {
	if c.client != nil {
		c.client.Close()
	}
}
