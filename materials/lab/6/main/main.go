package main

// main.go HAS FOUR TODOS - TODO_1 - TODO_4

import (
	"log"
	"net/http"
	"scrape/scrape"

	"github.com/gorilla/mux"
)

//TODO_1: Logging right now just happens, create a global constant integer LOG_LEVEL
//TODO_1: When LOG_LEVEL = 0 DO NOT LOG anything
//TODO_1: When LOG_LEVEL = 1 LOG API details only
//TODO_1: When LOG_LEVEL = 2 LOG API details and file matches (e.g., everything)

func main() {

	if scrape.LOG_LEVEL > 1 {
		log.Println("starting API server")
	}
	//create a new router
	router := mux.NewRouter()
	if scrape.LOG_LEVEL > 1 {
		log.Println("creating routes")
	}
	//specify endpoints
	router.HandleFunc("/", scrape.MainPage).Methods("GET")

	router.HandleFunc("/api-status", scrape.APISTATUS).Methods("GET")

	router.HandleFunc("/indexer", scrape.IndexFiles).Methods("GET")
	router.HandleFunc("/search", scrape.FindFile).Methods("GET")
	//TODO_2 router.HandleFunc("/addsearch/{regex}", scrape.TODOREPLACE).Methods("GET")
	router.HandleFunc("/addsearch/{regex}", scrape.AddRegex).Methods("GET")
	//TODO_3 router.HandleFunc("/clear", scrape.TODOREPLACE).Methods("GET")
	router.HandleFunc("/clear", scrape.ClearRegex).Methods("GET")
	//TODO_4 router.HandleFunc("/reset", scrape.TODOREPLACE).Methods("GET")
	router.HandleFunc("/reset", scrape.ResetRegex).Methods("GET")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}
