package main

import (
	"flag"
	"kia-logix/cmd/app"
	"kia-logix/config"
	cronJobs "kia-logix/cron_jobs"
	jobsWorker "kia-logix/jobs_worker"
	"log"
	"os"
	"sync"

	HTTPServer "kia-logix/api/rest"
)

func main() {
	cfg := readConfig()

	appContainer, err := app.NewAppContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	var wg sync.WaitGroup
	wg.Add(3)

	// Start HTTP server goroutine with recover
	go func() {
		defer wg.Done()
		defer recoverFromPanic("HTTPServer.Run")

		HTTPServer.Run(cfg, appContainer)
	}()

	// Start jobs worker goroutine with recover
	go func() {
		defer wg.Done()
		defer recoverFromPanic("jobsWorker.Run")

		jobsWorker.Run(cfg, appContainer)
	}()

	// Start cron jobs goroutine with recover
	go func() {
		defer wg.Done()
		defer recoverFromPanic("cronJobs.Run")

		cronJobs.Run(cfg, appContainer)
	}()

	wg.Wait()
}

// recoverFromPanic handles any panic and logs the error
func recoverFromPanic(funcName string) {
	if r := recover(); r != nil {
		log.Printf("Recovered from panic in %s: %v", funcName, r)
	}
}

var configPath = flag.String("config", "", "configuration path")

func readConfig() config.Config {
	flag.Parse()

	if cfgPathEnv := os.Getenv("APP_CONFIG_PATH"); len(cfgPathEnv) > 0 {
		*configPath = cfgPathEnv
	}

	if len(*configPath) == 0 {
		log.Fatal("configuration does not provide")
	}

	cfg := config.MustReadStandard(*configPath)

	return cfg
}
