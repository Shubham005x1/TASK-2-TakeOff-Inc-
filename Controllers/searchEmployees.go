package controllers

import (
    "TakeOff_Task-2/models"
    "encoding/json"
    "net/http"
    "strings"
	storage "TakeOff_Task-2/Storage"
)

func SearchEmployee(w http.ResponseWriter, r *http.Request) {
    // Extract query parameters from the request
    firstname := r.URL.Query().Get("firstname")
    lastname := r.URL.Query().Get("lastname")
    email := r.URL.Query().Get("email")
    role := r.URL.Query().Get("role")

    // Read all employees from storage
    employees, err := storage.ReadAllEmployees()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    var searchResults []models.Employee

    for _, employee := range employees {
        if (firstname == "" || strings.Contains(strings.ToLower(employee.FirstName), strings.ToLower(firstname))) &&
            (lastname == "" || strings.Contains(strings.ToLower(employee.LastName), strings.ToLower(lastname))) &&
            (email == "" || strings.Contains(strings.ToLower(employee.Email), strings.ToLower(email))) &&
            (role == "" || strings.Contains(strings.ToLower(employee.Role), strings.ToLower(role))) {
            searchResults = append(searchResults, employee)
        }
    }
	if len(searchResults) == 0 {
		w.Write([]byte("No Employee was found"))
	}
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(searchResults)
}
