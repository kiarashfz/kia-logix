package jobs_worker

import (
	worker "github.com/contribsys/faktory_worker_go"
	"kia-logix/cmd/app"
	"kia-logix/config"
)

func Run(cfg config.Config, appContainer *app.Container) {
	mgr := worker.NewManager()

	// Job registration
	mgr.Register("UpdateOrderStatus", appContainer.GetOrderService().UpdateOrderStatusTask)

	// use up to N goroutines to execute jobs
	mgr.Concurrency = 100

	// pull jobs from these queues, in this order of precedence
	mgr.ProcessStrictPriorityQueues("critical", "default", "low_priority")

	// Start processing jobs, this method does not return.
	mgr.Run()
}
