package client

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
)

type (
	Durable struct {
		c client.Client
	}
)

var (
	instance *Durable
	once     sync.Once
)

const DQueue = "D_QUEUE"

func New() *Durable {
	once.Do(func() {
		c, err := client.Dial(client.Options{})
		if err != nil {
			log.Fatalf("Failed to create Temporal client: %v", err)
		}

		instance = &Durable{c: c}
	})

	return instance
}

func (t *Durable) Client() client.Client {
	if t.c == nil {
		t = New()
	}

	return t.c
}

func (t *Durable) Close() {
	if t.c != nil {
		t.c.Close()
	}
}

func StartWorkflowOptions(queue string) client.StartWorkflowOptions {
	return client.StartWorkflowOptions{
		ID:        uuid.NewString(),
		TaskQueue: queue,
		RetryPolicy: &temporal.RetryPolicy{
			InitialInterval:    time.Second * 2,
			BackoffCoefficient: 2.0,
			MaximumAttempts:    5,
		},
	}
}

func retry() *temporal.RetryPolicy {
	return &temporal.RetryPolicy{
		InitialInterval:    time.Second,
		BackoffCoefficient: 2.0,
		MaximumInterval:    100 * time.Second,
		MaximumAttempts:    500,
	}
}

func ActivityOptions() workflow.ActivityOptions {
	return workflow.ActivityOptions{
		StartToCloseTimeout: time.Minute,
		RetryPolicy:         retry(),
	}
}
