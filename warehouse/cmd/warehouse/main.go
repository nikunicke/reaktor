package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/nikunicke/reaktor/warehouse"
	"github.com/nikunicke/reaktor/warehouse/api"
	"github.com/nikunicke/reaktor/warehouse/api/options"
	"github.com/nikunicke/reaktor/warehouse/http"
)

func main() {
	fmt.Println("warehouse_api started")

	client, err := api.NewClient(
		options.Client().ApplyURI("https://bad-api-assignment.reaktor.com/"))
	if err != nil {
		log.Fatal(err)
	}

	warehouse := warehouse.New(warehouse.ProductCategories...)
	warehouse.ProductService = api.NewProductService(client)
	warehouse.AvailabilityService = api.NewAvailabilityService(client)
	warehouse.UpdateInterval = time.NewTicker(time.Minute * 5)
	for {
		if err := warehouse.Update(); err != nil {
			fmt.Printf("Error: %v\n", err)
			fmt.Println("Retrying...")
		} else {
			break
		}
	}
	fmt.Println("init loop done")
	go func() {
		for {
			fmt.Println("auto update")
			select {
			case <-warehouse.UpdateInterval.C:
				warehouse.Update()
			}
		}
	}()

	server := http.NewServer()
	server.Addr = ":80"
	server.WarehouseService = warehouse
	server.ProductService = api.NewProductService(client)
	server.AvailabilityService = api.NewAvailabilityService(client)
	server.Open()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	server.Close()
	fmt.Println("\nwarehouse_api exiting")
}
