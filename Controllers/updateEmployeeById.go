package controllers

import (
	validations "TakeOff_Task-2/Validations"
	"TakeOff_Task-2/models"
	"TakeOff_Task-2/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// UpdateEmployeeByID updates an employee's details by ID.
func UpdateEmployeeById(w http.ResponseWriter, r *http.Request) {
	// Get the employee ID from the URL parameters.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Decode the request body to get the updated employee details.
	var updatedEmployee models.Employee
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Check if all required fields are provided.
	if updatedEmployee.ID == "" || updatedEmployee.FirstName == "" || updatedEmployee.LastName == "" || updatedEmployee.Email == "" || updatedEmployee.Role == "" || updatedEmployee.PhoneNo == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	// Validation for FirstName
	if validations.ValidNameEntry(updatedEmployee.FirstName) {
		http.Error(w, "Name Cannot contain Numbers please enter valid Name", http.StatusBadRequest)
		return
	}

	// Validation for LastName
	if validations.ValidNameEntry(updatedEmployee.LastName) {
		http.Error(w, "LastName Cannot contain Numbers please enter valid Name", http.StatusBadRequest)
		return
	}

	// Validation for Email
	err = validations.IsValidEmail(updatedEmployee.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Validation for PhoneNumber
	err = validations.IsNumberValid(updatedEmployee.PhoneNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Convert the employee ID to an integer for comparison.
	v, _ := strconv.Atoi(updatedEmployee.ID)
	if v == id {
		// If the provided ID matches the URL parameter, update the employee details.
		err = utils.UpdateEmployeeByID(id, updatedEmployee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		// If the provided ID doesn't match the URL parameter, return an error.
		http.Error(w, "Employee ID cannot be different", http.StatusBadRequest)
		return
	}

	// Respond with a success status and message.
	w.WriteHeader(http.StatusOK)
	val := fmt.Sprintf("Employee with ID %d is updated", id)
	w.Write([]byte(val))
}
