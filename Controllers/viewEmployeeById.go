package controllers

import (
	storage "TakeOff_Task-2/Storage"
	"TakeOff_Task-2/models"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func ViewEmployeeById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid employee Id", http.StatusBadRequest)
	}
	emp, err := storage.ReadAllEmployees()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var foundEmployee *models.Employee
	for _, employees := range emp {
		val, _ := strconv.Atoi(employees.ID)
		if val == id {
			foundEmployee = &employees
			break
		}

	}
	if foundEmployee == nil {
		http.Error(w, "Employee not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(foundEmployee)

}
