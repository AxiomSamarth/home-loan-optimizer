package main

import (
	"log"
	"net/http"
	"os"

	homeloan "github.com/AxiomSamarth/home-loan-optimizer/pkg/homeloan"
)

func handleRequests() {
	port := ":" + os.Getenv("PORT")
	http.HandleFunc("/", homeloan.Home)
	http.HandleFunc("/strategies", homeloan.Strategies)
	log.Fatal(http.ListenAndServe(port, nil))
}

func main() {
	handleRequests()
}
