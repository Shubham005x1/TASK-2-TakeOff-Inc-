package controllers

import (
	storage "TakeOff_Task-2/Storage"
	"TakeOff_Task-2/models"
	"encoding/csv"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

// DeleteEmployeeByID deletes an employee with a specific ID.
func DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {
	// Get the employee ID from the URL parameters.
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	// Handle the case where the employee ID is not a valid integer.
	if err != nil {
		http.Error(w, "Invalid Employee Id", http.StatusBadRequest)
	}
	// Read all employees from the storage.
	emp, err := storage.ReadAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var foundEmp []models.Employee
	var flag bool
	// Iterate through all employees to find the one to be deleted
	for _, employees := range emp {
		empid, _ := strconv.Atoi(employees.ID)
		if empid == id {
			flag = true

		} else {
			foundEmp = append(foundEmp, employees)
		}

	}
	// Handle the case where the employee with the specified ID was not found.
	if !flag {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	// Open the CSV file for writing.
	file, _ := os.OpenFile(storage.CsvFilePath, os.O_RDWR|os.O_TRUNC, 0666)

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	// Write the remaining employees (excluding the one to be deleted) to the CSV file.
	for _, val := range foundEmp {
		row := []string{
			val.ID,
			val.FirstName,
			val.LastName,
			val.Email,
			val.Password,
			val.PhoneNo,
			val.Role,
		}
		writer.Write(row)
	}
	val := fmt.Sprintf("Employee with ID %d is deleted", id)
	w.Write([]byte(val))

}
