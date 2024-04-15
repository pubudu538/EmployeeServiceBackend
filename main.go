package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type Employee struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Department string `json:"department"`
}

var employees = []Employee{
	{Id: "1234123", Name: "Mrs. Heily Feyers", Department: "IT"},
	{Id: "23451234", Name: "Mr. Brendon MacSmith", Department: "Sales"},
	{Id: "34561234", Name: "Mr. Peter Queenslander", Department: "IT"},
	{Id: "45671243", Name: "Miss. Liza MacAdams", Department: "Finance"},
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if jsonData, err := json.MarshalIndent(employees, "", "    "); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

func addEmployee(w http.ResponseWriter, r *http.Request) {
	var newEmployee Employee
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employees = append(employees, newEmployee)
	json.NewEncoder(w).Encode(newEmployee)
}

func editEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId := strings.TrimPrefix(r.URL.Path, "/employee/")
	for i, emp := range employees {
		if emp.Id == employeeId {
			var updatedEmployee Employee
			if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			employees[i] = updatedEmployee
			json.NewEncoder(w).Encode(updatedEmployee)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	employeeId := strings.TrimPrefix(r.URL.Path, "/employee/")
	for i, emp := range employees {
		if emp.Id == employeeId {
			employees = append(employees[:i], employees[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}

func main() {
	http.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getEmployees(w, r)
		case "POST":
			addEmployee(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			editEmployee(w, r)
		case "DELETE":
			deleteEmployee(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	http.ListenAndServe(":80", nil)
}
