package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

// declare database variables
const (
	DB_USER     = "postgres"
	DB_PASSWORD = "12345678"
	DB_NAME     = "movies"
)

// define struct Movie, which define the JSon fields going to be fetched
type Movie struct {
	MovieID   string `json:"movieid"`
	MovieName string `json:"moviename"`
}

// define struct JsonResponse, which will display the JSON response once the data is fetched
type JsonResponse struct {
	Type    string  `json:"type"`
	Data    []Movie `json:"data"`
	Message string  `json:"message"`
}

// // DB set up
func setupDB() *sql.DB {
	dbinfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASSWORD, DB_NAME)
	db, err := sql.Open("postgres", dbinfo)

	checkErr(err)

	return db
}

// home page
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome home!")
}

// Get all movies,
// response and request handlers
func GetMovies(w http.ResponseWriter, r *http.Request) {
	db := setupDB()

	printMessage("Getting movies ...")

	// get all movies from movies table that don't have movieID = "1"
	rows, err := db.Query("SELECT * FROM movies")

	// check erros
	checkErr(err)

	// var response [] JsonResponse
	var movies []Movie

	// for each movie
	for rows.Next() {
		var id int
		var movieID string
		var movieName string

		err = rows.Scan(&id, &movieID, &movieName)

		// check errors
		checkErr(err)

		movies = append(movies, Movie{MovieID: movieID, MovieName: movieName})
	}

	var response = JsonResponse{Type: "success", Data: movies}

	json.NewEncoder(w).Encode(response)

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// function for handling message
func printMessage(message string) {
	fmt.Println("")
	fmt.Println(message)
	fmt.Println("")
}

// main function
func main() {
	// Init the mux router
	router := mux.NewRouter()

	// Route handles and endpoints

	// homepage
	router.HandleFunc("/", homePage)

	// get all movies
	router.HandleFunc("/movies/", GetMovies).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))

}
