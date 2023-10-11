package controllers

import (
	storage "TakeOff_Task-2/Storage"
	"encoding/json"
	"net/http"
)

func ViewAllEmployees(w http.ResponseWriter, r *http.Request) {
	emp, _ := storage.ReadAllEmployees()
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(emp)
}
