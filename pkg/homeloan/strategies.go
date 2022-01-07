package homeloan

import (
	"fmt"
	"net/http"
)

// Home is route handler than returns the home page
func Home(w http.ResponseWriter, r *http.Request) {

	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}

// Strategies function accepts HTTP POST requests to compute
// recommendation for optimum home loan settlement strategies
func Strategies(w http.ResponseWriter, r *http.Request) {

	// TODO: Implementation is yet to done
}
