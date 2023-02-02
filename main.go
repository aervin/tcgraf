package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"main/tcgraf"
)

const (
	envFileLocation   = "/tcgraf/.env"
	grafanaKeyEnv     = "TC_GRAFANA_KEY"
	grafanaAddressEnv = "TC_GRAFANA_ADDRESS"
)

func main() {
	log.Print("tcgraf running...")

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	err := godotenv.Load(envFileLocation)
	if err != nil {
		log.Fatal("failed to load .env file")
	}
	address := os.Getenv(grafanaAddressEnv)
	if address == "" {
		log.Fatal("a grafana instance address must be provided")
	}
	apiKey := os.Getenv(grafanaKeyEnv)
	if apiKey == "" {
		log.Fatal("a grafana api key must be provided")
	}

	go func() {
		err := tcgraf.Start(address, apiKey)
		log.Print("server stopped unexpectedly", err.Error())
		os.Exit(1)
	}()
	log.Print("server started")

	<-shutdown
	log.Print("shutdown signal received")
	log.Print("exiting...")
	os.Exit(0)
}
