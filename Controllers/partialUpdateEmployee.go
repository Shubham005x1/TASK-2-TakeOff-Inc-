package controllers

import (
	storage "TakeOff_Task-2/Storage"
	validations "TakeOff_Task-2/Validations"
	"TakeOff_Task-2/models"
	"TakeOff_Task-2/utils"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

// PartialUpdateEmployee updates specific fields of an existing employee.
func PartialUpdateEmployee(w http.ResponseWriter, r *http.Request) {
	// Retrieve the employee ID from the URL path parameters
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}

	// Retrieve the existing employee from storage
	existingEmployee, err := storage.GetEmployeeByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var updatedEmployee models.Employee
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Update only the specified fields
	if updatedEmployee.FirstName != "" {
		if validations.ValidNameEntry(updatedEmployee.FirstName) {
			http.Error(w, "Name Cannot contain Numbers please enter valid Name", http.StatusBadRequest)
			return
		}
		existingEmployee.FirstName = updatedEmployee.FirstName
	}
	if updatedEmployee.LastName != "" {
		if validations.ValidNameEntry(updatedEmployee.LastName) {
			http.Error(w, "Name Cannot contain Numbers please enter valid Name", http.StatusBadRequest)
			return
		}
		existingEmployee.LastName = updatedEmployee.LastName
	}
	if updatedEmployee.Email != "" {
		err = validations.IsValidEmail(updatedEmployee.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		existingEmployee.Email = updatedEmployee.Email
	}
	if updatedEmployee.Role != "" {
		err = validations.IsValidRole(updatedEmployee.Email)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		existingEmployee.Role = updatedEmployee.Role
	}
	if updatedEmployee.PhoneNo != "" {
		err = validations.IsNumberValid(updatedEmployee.PhoneNo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		existingEmployee.PhoneNo = updatedEmployee.PhoneNo
	}

	// Perform the actual update operation
	err = utils.UpdateEmployeeByID(id, existingEmployee)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	val := fmt.Sprintf("Employee with ID %d is partially updated", id)
	w.Write([]byte(val))
}
