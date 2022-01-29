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

func init() {
	envChecks()
}

func main() {
	handleRequests()
}

func envChecks() {
	port, portExist := os.LookupEnv("PORT")

	if !portExist || port == "" {
		log.Fatal("PORT must be set in .env and not empty")
	}
}
