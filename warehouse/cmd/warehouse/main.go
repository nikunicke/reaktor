package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"time"

	"github.com/nikunicke/reaktor/warehouse"
	"github.com/nikunicke/reaktor/warehouse/api"
	"github.com/nikunicke/reaktor/warehouse/api/options"
	"github.com/nikunicke/reaktor/warehouse/http"
)

var mu sync.RWMutex

func main() {
	fmt.Println("warehouse_api started")

	client, err := api.NewClient(
		options.Client().ApplyURI("https://bad-api-assignment.reaktor.com/"))
	if err != nil {
		log.Fatal(err)
	}

	warehouseService := warehouse.New(warehouse.ProductCategories...)
	warehouseService.ProductService = api.NewProductService(client)
	warehouseService.AvailabilityService = api.NewAvailabilityService(client)
	warehouseService.UpdateInterval = time.NewTicker(time.Minute * 5)
	for {
		if err := warehouseService.Update(); err != nil {
			fmt.Printf("Error: %v\n", err)
			fmt.Println("Retrying...")
		} else {
			break
		}
	}
	go func() {
		for {
			select {
			case <-warehouseService.UpdateInterval.C:
				{
					mu.Lock()
					warehouseService.Update()
					mu.Unlock()
				}
			}
		}
	}()

	port := os.Getenv("PORT")
	if port == "" {
		port = ":5000"
	}
	server := http.NewServer()
	server.Addr = ":" + port
	server.WarehouseService = warehouseService
	server.ProductService = warehouseService
	server.Open()
	fmt.Println("Server running on port:", server.URL())

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	server.Close()
	fmt.Println("\nwarehouse_api exiting")
}
