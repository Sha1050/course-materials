package classes

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type student struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	FatherName string `json:"father_name"`
	Class      string `json:"class"`
	Section    string `json:"section"`
}

var students []student

func InitStudent() {
	var initStudent student
	initStudent.Id = "1"
	initStudent.Name = "First Student"
	initStudent.FatherName = "First student father name"
	initStudent.Class = "O Level"
	initStudent.Section = "C"

	students = append(students, initStudent)
}

// Get All Students Information
func GetStudents(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	// Get All Students
	json.NewEncoder(w).Encode(students)
}

// Get Specific Student Information
func GetStudent(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	for _, std := range students {
		if std.Id == params["id"] {
			json.NewEncoder(w).Encode(std)
			break
		}
	}
	//TODO : Provide a response if there is no such job
	//w.Write(jsonResponse)
}

// Create New Student
func CreateNewStudent(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var newStudent student
	json.NewDecoder(r.Body).Decode(&newStudent)
	newStudent.Id = strconv.Itoa(len(students) + 1)
	students = append(students, newStudent)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newStudent)

}

// Delete Student
func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s DELETE end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/txt")
	w.WriteHeader(http.StatusOK)
	params := mux.Vars(r)

	response := make(map[string]string)

	response["status"] = "No Such ID to Delete"
	for index, std := range students {
		if std.Id == params["id"] {
			students = append(students[:index], students[index+1:]...)
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

// Update Specific Student
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	log.Printf("Entering %s end point", r.URL.Path)
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	for index, std := range students {
		if std.Id == params["id"] {
			students = append(students[:index], students[index+1:]...)
			var newStd student
			json.NewDecoder(r.Body).Decode(&newStd)
			newStd.Id = params["id"]
			students = append(students, newStd)
			json.NewEncoder(w).Encode(newStd)
			return
		}
	}
}
