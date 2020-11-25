package main

import (
	"fmt"
	"os"
	"os/signal"

	"github.com/nikunicke/reaktor/warehouse_api/http"
)

func main() {
	fmt.Println("warehouse_api started")
	server := http.NewServer()
	server.Addr = ":80"
	server.Open()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	server.Close()
	fmt.Println("\nwarehouse_api exiting")
}
