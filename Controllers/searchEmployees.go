package controllers

import (
	storage "TakeOff_Task-2/Storage"
	"TakeOff_Task-2/models"
	"encoding/json"
	"net/http"
	"strings"
)

// SearchEmployee searches for employees based on a query parameter.
func SearchEmployee(w http.ResponseWriter, r *http.Request) {
	// Get the search query from the URL parameters.
	query := r.URL.Query().Get("search")

	// Handle the case where no search query is provided.
	if query == "" {
		http.Error(w, "Search query not provided (Query should be named search)", http.StatusBadRequest)
		return
	}

	// Retrieve all employees from the storage.
	employees, err := storage.ReadAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Initialize an empty slice to hold the search results.
	var searchResults []models.Employee

	// Iterate through all employees to find matches based on the query.
	for _, employee := range employees {
		if strings.Contains(strings.ToLower(employee.FirstName), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(employee.LastName), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(employee.Email), strings.ToLower(query)) ||
			strings.Contains(strings.ToLower(employee.Role), strings.ToLower(query)) {
			searchResults = append(searchResults, employee)
		}
	}

	// If no matching employees were found, provide a response indicating this.
	if len(searchResults) == 0 {
		w.Write([]byte("No Employee was found"))
	}

	// Set the response content type to JSON and encode the search results.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResults)
}
