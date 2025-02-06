package client

import (
	"log"
	"sync"

	"go.temporal.io/sdk/client"
)

type (
	TemporalClient struct {
		c client.Client
	}
)

var (
	instance *TemporalClient
	once     sync.Once
)

const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

func New() *TemporalClient {
	once.Do(func() {
		c, err := client.Dial(client.Options{})
		if err != nil {
			log.Fatalf("Failed to create Temporal client: %v", err)
		}

		instance = &TemporalClient{c: c}
	})

	return instance
}

func (t *TemporalClient) Client() client.Client {
	if t.c == nil {
		t = New()
	}

	return t.c
}

func (t *TemporalClient) Close() {
	if t.c != nil {
		t.c.Close()
	}
}
