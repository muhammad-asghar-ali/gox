package worker

import (
	c "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/muhammad-asghar-ali/gox/fintech/internal/services/transactions"
	"github.com/muhammad-asghar-ali/gox/fintech/temporal/client"
)

type (
	WorkerConfig struct {
		Client c.Client
		Queue  string
	}
)

func NewWorker(c c.Client, queue string) worker.Worker {
	wf := worker.New(c, queue, worker.Options{})

	w := transactions.Workflow{}
	a := transactions.Activities{}

	wf.RegisterWorkflow(w.Transaction)
	wf.RegisterActivity(a.CheckBalance)
	wf.RegisterActivity(a.CreateTransaction)
	wf.RegisterActivity(a.MatchAccount)
	wf.RegisterActivity(a.Tranfer)
	wf.RegisterActivity(a.GetAccount)

	return wf
}

func Run() error {
	queue := client.MoneyTransferTaskQueueName

	temporalClient := client.New()
	defer temporalClient.Close()

	w := NewWorker(temporalClient.Client(), queue)
	return w.Run(worker.InterruptCh())
}
