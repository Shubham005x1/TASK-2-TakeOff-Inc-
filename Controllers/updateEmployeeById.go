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

func UpdateEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee ID", http.StatusBadRequest)
		return
	}
	var updatedEmployee models.Employee
	err = json.NewDecoder(r.Body).Decode(&updatedEmployee)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if updatedEmployee.ID == "" || updatedEmployee.FirstName == "" || updatedEmployee.LastName == "" || updatedEmployee.Email == "" || updatedEmployee.Role == "" || updatedEmployee.PhoneNo == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}

	//Validation for FirstName
	if validations.ValidNameEntry(updatedEmployee.FirstName) {
		http.Error(w, "Name Cannot contain Numbers please enter valid Name", http.StatusBadRequest)
		return
	}
	//Validation for LastName
	if validations.ValidNameEntry(updatedEmployee.LastName) {
		http.Error(w, "LastName Cannot contain Numbers please enter valid Name", http.StatusBadRequest)
		return
	}
	//Validation for Email
	err = validations.IsValidEmail(updatedEmployee.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Validation for PhoneNumber
	err = validations.IsNumberValid(updatedEmployee.PhoneNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	v, _ := strconv.Atoi(updatedEmployee.ID)
	if v == id {
		err = utils.UpdateEmployeeByID(id, updatedEmployee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	} else {
		http.Error(w, "Employee ID cannot be different", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusOK)
	val := fmt.Sprintf("Employee with ID %d is updated", id)
	w.Write([]byte(val))

}
