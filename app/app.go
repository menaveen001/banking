package app

import (
	"crud/handllers"
	"log"
	"net/http"
)

func Start() {

	// define routes
	http.HandleFunc("/greet", handllers.Greet)
	http.HandleFunc("/customer", handllers.GetAllCustomers)

	//statring server
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
