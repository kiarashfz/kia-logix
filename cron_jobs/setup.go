package cronJobs

import (
	"fmt"
	faktory "github.com/contribsys/faktory/client"
	"github.com/go-co-op/gocron/v2"
	"golang.org/x/net/context"
	"kia-logix/cmd/app"
	"kia-logix/config"
	"kia-logix/internal/orders"
	"kia-logix/service"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func pushUpdateNeedOrdersToQueue(faktoryClient *faktory.Client, ordersSlice []orders.Order) {
	for _, order := range ordersSlice {
		if order.Provider.Name == "podro" {
			job := faktory.NewJob("UpdateOrderStatus", order)
			job.Queue = "critical"
			err := faktoryClient.Push(job)
			if err != nil {
				fmt.Println("Error pushing job")
			}
		}
	}
}

func getUpdateNeedOrders(ctx context.Context, orderService service.IOrderService) []orders.Order {
	orderObjs, err := orderService.GetUpdateNeedOrders(ctx)
	if err != nil {
		log.Println(err)
		return nil
	}
	return orderObjs
}

func Run(cfg config.Config, appContainer *app.Container) {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}

	// add a Daily job to the scheduler
	s.NewJob(
		gocron.DailyJob(
			1, // Runs every day
			gocron.NewAtTimes(
				gocron.NewAtTime(uint(cfg.CronJobs.UpdateOrdersStatusJob.Hour), uint(cfg.CronJobs.UpdateOrdersStatusJob.Minute), 00),
			),
		),
		gocron.NewTask(
			func() {
				client, err := faktory.Open()
				if err != nil {
					panic(err)
				}
				ctx := context.Background()
				orderObjs := getUpdateNeedOrders(ctx, appContainer.GetOrderService())
				pushUpdateNeedOrdersToQueue(client, orderObjs)
			},
		),
		gocron.WithName("Daily Order Update"),
	)

	// Start the scheduler
	s.Start()

	// Set up a channel to listen for interrupt signals
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-sigChan
		log.Println("\nInterrupt signal received. Exiting...")
		_ = s.Shutdown()
		os.Exit(0)
	}()

	select {}
}
