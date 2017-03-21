package main

import (
	"flag"
	"log"

	"code.cloudfoundry.org/lager"

	"github.com/kubernetes-incubator/service-catalog/contrib/pkg/broker/server"

	"github.com/deis/minio-broker/controller"
)

var (
	configFilePath string
	port           string

	logLevels = map[string]lager.LogLevel{
		"DEBUG": lager.DEBUG,
		"INFO":  lager.INFO,
		"ERROR": lager.ERROR,
		"FATAL": lager.FATAL,
	}
)

func init() {
	flag.StringVar(&configFilePath, "config", "", "Location of the config file")
	flag.StringVar(&port, "port", "3000", "Listen port")
}

func main() {
	log.Println("Starting Minio broker...")

	flag.Parse()

	config, err := LoadConfig(configFilePath)
	if err != nil {
		log.Fatalf("Error loading config file: %s", err)
	}

	controller, err := controller.CreateController(config.MinioConfig)
	if err != nil {
		log.Fatal(err)
	}
	server.Start(8080, controller)
}
