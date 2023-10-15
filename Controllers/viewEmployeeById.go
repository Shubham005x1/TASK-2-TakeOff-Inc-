package controllers

import (
	storage "TakeOff_Task-2/Storage"
	"TakeOff_Task-2/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// ViewEmployeeByID retrieves and returns an employee by their ID.
func ViewEmployeeById(w http.ResponseWriter, r *http.Request) {
	// Get the employee ID from the URL parameters.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Read all employees from the storage.
	emp, err := storage.ReadAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundEmployee *models.Employee

	// Iterate through all employees to find the one with the specified ID.
	for _, employee := range emp {
		val, _ := strconv.Atoi(employee.ID)
		if val == id {
			// Set the found employee.
			foundEmployee = &employee
			break
		}
	}

	// If no employee with the specified ID was found, return an error.
	if foundEmployee == nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	// Set the response content type to JSON and encode the employee as JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundEmployee)
}
