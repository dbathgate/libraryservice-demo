package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"./config"
	"./metrics"
	"./route"
)

func main() {
	configFilePath := flag.String("config", "config/config.yml", "config file")
	version := flag.String("version", "v1", "app version")

	flag.Parse()

	source, err := ioutil.ReadFile(*configFilePath)
	if err != nil {
		panic(err)
	}

	config.LoadConfig(source)
	config.Version = *version

	metrics.InitMetrics()

	log.Printf("Starting library service on port %d...", config.Config.Server.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", config.Config.Server.Port), route.SetupHttpRoutes()))
}
