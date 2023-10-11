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
	router.HandleFunc("/create", controllers.CreateEmployee).Methods("POST")
	router.HandleFunc("/viewAll", controllers.ViewAllEmployees).Methods("GET")
	router.HandleFunc("/view/{id}", controllers.ViewEmployeeById).Methods("GET")
	router.HandleFunc("/delete/{id}", controllers.DeleteEmployeeById).Methods("POST")
	router.HandleFunc("/update/{id}", controllers.UpdateEmployeeById).Methods("POST")
	log.Println("Server started on :4041")
	log.Fatal(http.ListenAndServe(":4041", router))
}
