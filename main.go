package main

import (
	"encoding/json"
	"log"
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
	// Create ServeMux instances for port 80 and port 8080
	mux80 := http.NewServeMux()
	mux8080 := http.NewServeMux()

	// Handlers for port 80
	mux80.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getEmployees(w, r)
		case "POST":
			addEmployee(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux80.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			editEmployee(w, r)
		case "DELETE":
			deleteEmployee(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Handlers for port 8080
	mux8080.HandleFunc("/employee", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "GET":
			getEmployeesV2(w, r)
		case "POST":
			addEmployeeV2(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})
	mux8080.HandleFunc("/employee/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case "PUT":
			editEmployeeV2(w, r)
		case "DELETE":
			deleteEmployeeV2(w, r)
		default:
			w.WriteHeader(http.StatusMethodNotAllowed)
		}
	})

	// Run HTTP servers on ports 80 and 8080 concurrently
	go func() {
		if err := http.ListenAndServe(":80", mux80); err != nil {
			log.Fatalf("Server on port 80 failed to start: %v", err)
		}
	}()
	go func() {
		if err := http.ListenAndServe(":8080", mux8080); err != nil {
			log.Fatalf("Server on port 8080 failed to start: %v", err)
		}
	}()

	// Wait indefinitely
	select {}
}
