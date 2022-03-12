package main

import (
	"log"
	"net/http"
	"wyoassign/classes"
	"wyoassign/wyoassign"

	"github.com/gorilla/mux"
)

func main() {
	wyoassign.InitAssignments()
	classes.InitStudent()
	log.Println("starting API server")
	//create a new router
	router := mux.NewRouter()
	log.Println("creating routes")
	//specify endpoints
	router.HandleFunc("/api-status", wyoassign.APISTATUS).Methods("GET")
	router.HandleFunc("/assignments", wyoassign.GetAssignments).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.Getassigment).Methods("GET")
	router.HandleFunc("/assignment/{id}", wyoassign.Deleteassignment).Methods("DELETE")
	router.HandleFunc("/assignment", wyoassign.Deleteassignment).Methods("POST")
	router.HandleFunc("/assignment/{id}", wyoassign.Updateassignment).Methods("PUT")

	// Student Information
	router.HandleFunc("/students", classes.GetStudents).Methods("GET")
	router.HandleFunc("/student/{id}", classes.GetStudent).Methods("GET")
	router.HandleFunc("/student/{id}", classes.DeleteStudent).Methods("DELETE")
	router.HandleFunc("/student", classes.CreateNewStudent).Methods("POST")
	router.HandleFunc("/student/{id}", classes.UpdateStudent).Methods("PUT")

	http.Handle("/", router)

	//start and listen to requests
	http.ListenAndServe(":8080", router)

}
