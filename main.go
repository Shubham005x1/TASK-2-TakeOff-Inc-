package main

import (
	controllers "TakeOff_Task-2/Controllers"
	storage "TakeOff_Task-2/Storage"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	storage.InitializeCSV()
	storage.WriteHeadersToCSV()
	router := mux.NewRouter()

	router.HandleFunc("/employees", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/employees", controllers.ViewAllEmployees).Methods("GET")
	router.HandleFunc("/employees/{id}", controllers.ViewEmployeeById).Methods("GET")
	router.HandleFunc("/employees/{id}", controllers.DeleteEmployeeById).Methods("DELETE")
	router.HandleFunc("/employees/{id}", controllers.UpdateEmployeeById).Methods("PUT")
	router.HandleFunc("/search", controllers.SearchEmployee).Methods("GET")

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
