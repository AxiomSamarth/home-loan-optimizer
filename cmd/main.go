package main

import (
	"log"
	"net/http"

	homeloan "github.com/AxiomSamarth/home-loan-optimizer/pkg/homeloan"
)

func handleRequests() {
	http.HandleFunc("/", homeloan.Home)
	http.HandleFunc("/strategies", homeloan.Strategies)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func main() {
	handleRequests()
}
