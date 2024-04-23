package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

type EmployeeV2 struct {
	Employee
	Gender string `json:"gender"`
}

var employeesV2 = []EmployeeV2{
	{
		Employee: Employee{
			Id:         "1234123",
			Name:       "Mrs. Heily Feyers",
			Department: "IT",
		},
		Gender: "Female",
	},
	{
		Employee: Employee{
			Id:         "23451234",
			Name:       "Mr. Brendon MacSmith",
			Department: "Sales",
		},
		Gender: "Male",
	},
	{
		Employee: Employee{
			Id:         "34561234",
			Name:       "Mr. Peter Queenslander",
			Department: "IT",
		},
		Gender: "Male",
	},
	{
		Employee: Employee{
			Id:         "45671243",
			Name:       "Miss. Liza MacAdams",
			Department: "Finance",
		},
		Gender: "Female",
	},
}

func getEmployeesV2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if jsonData, err := json.MarshalIndent(employeesV2, "", "    "); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	} else {
		w.Write(jsonData)
	}
}

func addEmployeeV2(w http.ResponseWriter, r *http.Request) {
	var newEmployee EmployeeV2
	if err := json.NewDecoder(r.Body).Decode(&newEmployee); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	employeesV2 = append(employeesV2, newEmployee)
	json.NewEncoder(w).Encode(newEmployee)
}

func editEmployeeV2(w http.ResponseWriter, r *http.Request) {
	employeeId := strings.TrimPrefix(r.URL.Path, "/employee/")
	for i, emp := range employeesV2 {
		if emp.Id == employeeId {
			var updatedEmployee EmployeeV2
			if err := json.NewDecoder(r.Body).Decode(&updatedEmployee); err != nil {
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			employeesV2[i] = updatedEmployee
			json.NewEncoder(w).Encode(updatedEmployee)
			return
		}
	}
	http.NotFound(w, r)
}

func deleteEmployeeV2(w http.ResponseWriter, r *http.Request) {
	employeeId := strings.TrimPrefix(r.URL.Path, "/employee/")
	for i, emp := range employeesV2 {
		if emp.Id == employeeId {
			employeesV2 = append(employeesV2[:i], employeesV2[i+1:]...)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}
	http.NotFound(w, r)
}
