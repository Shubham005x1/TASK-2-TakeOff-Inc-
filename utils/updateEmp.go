package utils

import (
	storage "TakeOff_Task-2/Storage"
	"TakeOff_Task-2/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func UpdateEmployeeByID(id int, updatedEmployee models.Employee) error {
	employees, err := storage.ReadAllEmployees() //Reading all data from CSV file
	if err != nil {
		return err
	}

	var updated bool

	for i, employee := range employees { // Parsing through employees data from CSV
		empid, _ := strconv.Atoi(employee.ID)
		if empid == id {
			employees[i] = updatedEmployee
			updated = true
			break
		}
	}

	if !updated {
		return fmt.Errorf("Employee with ID %v not found", id)
	}
	file, err := os.OpenFile(storage.CsvFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	for i := 0; i < len(employees); i++ {
		row := []string{
			employees[i].ID,
			employees[i].FirstName,
			employees[i].LastName,
			employees[i].Email,
			employees[i].Password,
			employees[i].PhoneNo,
			employees[i].Role,
		}
		err = writer.Write(row)
	}

	if err != nil {
		return err
	}
	return nil
}
