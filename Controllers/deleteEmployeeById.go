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

func DeleteEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid Employee Id", http.StatusBadRequest)
	}
	emp, err := storage.ReadAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	var foundEmp []models.Employee
	var flag bool
	for _, employees := range emp {
		empid, _ := strconv.Atoi(employees.ID)
		if empid == id {
			flag = true

		} else {
			foundEmp = append(foundEmp, employees)
		}

	}
	if !flag {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}

	file, _ := os.OpenFile(storage.CsvFilePath, os.O_RDWR|os.O_TRUNC, 0666)

	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
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
