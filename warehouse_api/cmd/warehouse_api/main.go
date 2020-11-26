package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/nikunicke/reaktor/warehouse_api/bad_api"
	"github.com/nikunicke/reaktor/warehouse_api/bad_api/options"
	"github.com/nikunicke/reaktor/warehouse_api/http"
)

func main() {
	fmt.Println("warehouse_api started")
	client, err := bad_api.NewClient(
		options.Client().ApplyURI("https://bad-api-assignment.reaktor.com/"))
	if err != nil {
		log.Fatal(err)
	}
	server := http.NewServer()
	server.Addr = ":80"
	server.ProductService = bad_api.NewProductService(client)
	server.AvailabilityService = bad_api.NewAvailabilityService(client)
	server.Open()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	server.Close()
	fmt.Println("\nwarehouse_api exiting")
}
