package main

import (
	"fmt"
	"log"
	"net/http"
)

func handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello my friiiiend")
}

func main() {
	fmt.Println("Hello world")

	http.HandleFunc("/", handle)
	log.Fatal(http.ListenAndServe(":80", nil))
}
