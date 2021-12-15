package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// main function
func main() {
	// Init the mux router
	router := mux.NewRouter()

	// Route handles and endpoints

	// homepage
	router.HandleFunc("/", homePage)

	log.Fatal(http.ListenAndServe(":8080", router))

}
