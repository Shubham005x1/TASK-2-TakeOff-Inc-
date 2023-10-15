package controllers

import (
	storage "TakeOff_Task-2/Storage"
	"encoding/json"
	"net/http"
)

// ViewAllEmployees retrieves and returns all employees from the storage.
func ViewAllEmployees(w http.ResponseWriter, r *http.Request) {
	// Read all employees from the storage.
	emp, _ := storage.ReadAllEmployees()

	// Set the response content type to JSON.
	w.Header().Set("Content-Type", "application/json")

	// Encode and send the employees as JSON response.
	json.NewEncoder(w).Encode(emp)
}
