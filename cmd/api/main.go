package main

import (
	"flag"
	"kia-logix/cmd/api/app"
	"kia-logix/config"
	"log"
	"os"

	HTTPServer "kia-logix/api/rest"
)

func main() {
	cfg := readConfig()

	appContainer, err := app.NewAppContainer(cfg)
	if err != nil {
		log.Fatal(err)
	}

	HTTPServer.Run(cfg, appContainer)
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
