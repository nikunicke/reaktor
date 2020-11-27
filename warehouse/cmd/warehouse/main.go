package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

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

	exec := warehouse.JobExecutor{
		ProductService:      api.NewProductService(client),
		AvailabilityService: api.NewAvailabilityService(client),
	}
	job := warehouse.Job{Type: warehouse.JobTypeGetAvailability}
	exec.ExecuteJob(&job)
	// c1 := make(chan string)
	// c2 := make(chan string)
	// c3 := make(chan string)
	// c4 := make(chan string)
	// c5 := make(chan string)
	// c6 := make(chan string)

	// go func() {
	// 	client.Get(warehouse.Jackets)
	// 	for i := 0; i < 4; i++ {
	// 		fmt.Println("processing availability:", <-c6)
	// 	}
	// 	c1 <- "Got jackets"
	// }()

	// go func() {
	// 	client.Get("availability/derp")
	// 	c6 <- "derp"
	// 	c2 <- "Got availability/derp"
	// }()

	// go func() {
	// 	client.Get("availability/reps")
	// 	c6 <- "reps"
	// 	c3 <- "Got availability/reps"
	// }()

	// go func() {
	// 	client.Get("availability/xoon")
	// 	c6 <- "xoon"
	// 	c4 <- "Got availability/xoon"
	// }()

	// go func() {
	// 	client.Get("availability/nouke")
	// 	c6 <- "nouke"
	// 	c5 <- "Got availability/nouke"
	// }()

	// for i := 0; i < 5; i++ {
	// 	select {
	// 	case msg1 := <-c1:
	// 		fmt.Println("recieved:", msg1)
	// 	case msg2 := <-c2:
	// 		fmt.Println("received:", msg2)
	// 	case msg3 := <-c3:
	// 		fmt.Println("received:", msg3)
	// 	case msg4 := <-c4:
	// 		fmt.Println("received:", msg4)
	// 	case msg5 := <-c5:
	// 		fmt.Println("received:", msg5)
	// 	}
	// }
	fmt.Println("done")
	server := http.NewServer()
	server.Addr = ":80"
	server.ProductService = api.NewProductService(client)
	server.AvailabilityService = api.NewAvailabilityService(client)
	server.Open()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	server.Close()
	fmt.Println("\nwarehouse_api exiting")
}
