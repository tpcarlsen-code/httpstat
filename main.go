package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/tpcarlsen-code/mon2http/config"
	"github.com/tpcarlsen-code/mon2http/entities"
	"github.com/tpcarlsen-code/mon2http/http"
	"github.com/tpcarlsen-code/mon2http/monitors"
	"github.com/tpcarlsen-code/mon2http/storage"
)

func main() {
	cnfPath := flag.String("c", "", "optional path to config file")
	flag.Parse()
	_ = cnfPath
	//config.LoadConfig(*cnfPath)
	log.Printf("Starting up using HTTP port: %d\n", config.Port)
	statusStorage := &storage.Status{}
	valuesStorage := &storage.Values{}
	server := http.NewServer(config.Port, config.AccessToken, statusStorage, valuesStorage)

	ec := make(chan error)
	// Start HTTP server.
	go func() {
		ec <- server.Start()
	}()

	manager := monitors.NewManager()
	manager.Init(config.NumSamples)

	ticker := time.NewTicker(config.UpdateInterval).C
	var alerts []monitors.Alert
	var values []entities.Value
	var status entities.Status
	for {
		status.Status = entities.StatusOk
		status.Alerts = []string{}

		alerts, values = manager.Update()
		if len(alerts) > 0 {
			status.Status = entities.StatusAlert
			for _, a := range alerts {
				status.Alerts = append(status.Alerts, string(a))
			}
		}
		statusStorage.Set(status)
		valuesStorage.Set(values)
		select {
		case err := <-ec:
			{
				log.Printf("ERROR: %s\n", err.Error())
				os.Exit(1)
			}
		case <-ticker:
		}
	}
}
