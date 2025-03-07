package wyoassign

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Response struct {
	Assignments []Assignment `json:"assignments"`
}

type Assignment struct {
	Id          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"desc"`
	Points      int    `json:"points"`
}

var Assignments []Assignment

const Valkey string = "FooKey"

func InitAssignments() {
	var assignment Assignment
	assignment.Id = "Mike1A"
	assignment.Title = "lab 4"
	assignment.Description = "some lab this guy made yesterday"
	assignment.Points = 20
	Assignments = append(Assignments, assignment)
}

func APISTATUS(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "API is up and running")
}

func GetAssignments(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	var response Response

	response.Assignments = Assignments

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	jsonResponse, err := json.Marshal(response)

	if err != nil {
		return
	}

	//TODO
	w.Write(jsonResponse)
}

func Getassigment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, assignment := range Assignments {
		if assignment.Id == params["id"] {
			json.NewEncoder(w).Encode(assignment)
			break
		}
	}
	//TODO : Provide a response if there is no such job
	//w.Write(jsonResponse)
}

func Deleteassignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, assignment := range Assignments {
		if assignment.Id == params["id"] {
			Assignments = append(Assignments[:index], Assignments[index+1:]...)
			response["status"] = "Success"
			break
		}
	}

	jsonResponse, err := json.Marshal(response)
	if err != nil {
		return
	}
	w.Write(jsonResponse)
}

func Updateassignment(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	params := mux.Vars(r)
	r.ParseForm()

	for index, assignment := range Assignments {
		if assignment.Id == params["id"] {
			Assignments = append(Assignments[:index], Assignments[index+1:]...)
			var assignment Assignment
			assignment.Id = params["id"]
			assignment.Title = r.FormValue("title")
			assignment.Description = r.FormValue("desc")
			assignment.Points, _ = strconv.Atoi(r.FormValue("points"))
			Assignments = append(Assignments, assignment)
			json.NewEncoder(w).Encode(assignment)
			return
		}
	}
}

func Createassignment(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var assignment Assignment
	r.ParseForm()
	// Possible TODO: Better Error Checking!
	// Possible TODO: Better Logging
	if r.FormValue("id") != "" {
		assignment.Id = r.FormValue("id")
		assignment.Title = r.FormValue("title")
		assignment.Description = r.FormValue("desc")
		assignment.Points, _ = strconv.Atoi(r.FormValue("points"))
		Assignments = append(Assignments, assignment)
		w.WriteHeader(http.StatusCreated)
	}
	w.WriteHeader(http.StatusNotFound)

}
