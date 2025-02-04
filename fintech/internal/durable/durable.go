package durable

import (
	"log"
	"sync"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"
)

type (
	Temporal struct {
		client client.Client
	}
)

var (
	instance *Temporal
	once     sync.Once
)

const MoneyTransferTaskQueueName = "TRANSFER_MONEY_TASK_QUEUE"

func Set() *Temporal {
	once.Do(func() {
		c, err := client.Dial(client.Options{})
		if err != nil {
			log.Fatalf("Failed to create Temporal client: %v", err)
		}
		instance = &Temporal{client: c}
	})

	return instance
}

func (t *Temporal) Client() client.Client {
	if t.client == nil {
		t = Set()
	}

	return t.client
}

func (t *Temporal) Worker(queue string) worker.Worker {
	return worker.New(t.client, queue, worker.Options{})
}

func Run() error {
	t := Set()
	defer t.client.Close()

	w := t.Worker(MoneyTransferTaskQueueName)
	return w.Run(worker.InterruptCh())
}
