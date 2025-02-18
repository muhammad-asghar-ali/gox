package worker

import (
	c "go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"github.com/muhammad-asghar-ali/gox/fintech/durable/client"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/services/transactions"
	"github.com/muhammad-asghar-ali/gox/fintech/internal/services/users"
)

type (
	WorkerConfig struct {
		Client c.Client
		Queue  string
	}
)

func NewWorker(c c.Client, queue string) worker.Worker {
	wf := worker.New(c, queue, worker.Options{})

	// ----- transaction workflows and activities
	tw := transactions.Workflow{}
	ta := transactions.Activities{}
	wf.RegisterWorkflow(tw.Transaction)
	wf.RegisterActivity(ta.CheckBalance)
	wf.RegisterActivity(ta.CreateTransaction)
	wf.RegisterActivity(ta.MatchAccount)
	wf.RegisterActivity(ta.Tranfer)
	wf.RegisterActivity(ta.GetAccount)

	// ----- user workflows and activities
	uw := users.Workflow{}
	ua := users.Activities{}
	wf.RegisterWorkflow(uw.Login)
	wf.RegisterWorkflow(uw.Register)

	wf.RegisterActivity(ua.CheckUser)
	wf.RegisterActivity(ua.GenerateToken)
	wf.RegisterActivity(ua.GetUserByID)
	wf.RegisterActivity(ua.CreateUser)
	wf.RegisterActivity(ua.CreateUserAccount)

	return wf
}

func Run() error {
	queue := client.DQueue

	temporalClient := client.New()
	defer temporalClient.Close()

	w := NewWorker(temporalClient.Client(), queue)
	return w.Run(worker.InterruptCh())
}
