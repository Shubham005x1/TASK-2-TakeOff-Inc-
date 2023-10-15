package controllers

import (
	storage "TakeOff_Task-2/Storage"
	validations "TakeOff_Task-2/Validations"
	"TakeOff_Task-2/models"
	"TakeOff_Task-2/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var emp models.Employee

	err := json.NewDecoder(r.Body).Decode(&emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Check if fields are empty
	if emp.ID == "" || emp.FirstName == "" || emp.LastName == "" || emp.Email == "" || emp.Role == "" || emp.PhoneNo == "" {
		http.Error(w, "All fields are required", http.StatusBadRequest)
		return
	}
	// Check for duplicate ID
	v, _ := strconv.Atoi(emp.ID)
	if utils.IsDuplicateID(v) {
		http.Error(w, "Employee ID already exists", http.StatusBadRequest)
		return
	}
	//Validation for FirstName
	if validations.ValidNameEntry(emp.FirstName) {
		http.Error(w, "Name Cannot contain Numbers please enter valid Name", http.StatusBadRequest)
		return
	}
	//Validation for LastName
	if validations.ValidNameEntry(emp.LastName) {
		http.Error(w, "LastName Cannot contain Numbers please enter valid Name", http.StatusBadRequest)
		return
	}
	//Validation for Email
	err = validations.IsValidEmail(emp.Email)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	//Validation for PhoneNumber
	err = validations.IsNumberValid(emp.PhoneNo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = storage.WriteEmployeeToCSV(emp)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)
	val := fmt.Sprintf("Employee with ID %s is created", emp.ID)
	w.Write([]byte(val))
}
