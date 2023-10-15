package utils

import (
	storage "TakeOff_Task-2/Storage"
	"TakeOff_Task-2/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// UpdateEmployeeByID updates an employee's details by ID.
func UpdateEmployeeByID(id int, updatedEmployee models.Employee) error {
	employees, err := storage.ReadAllEmployees() // Read all employee data from CSV file.
	if err != nil {
		return err
	}

	var updated bool

	// Iterate through employees to find the one with the specified ID.
	for i, employee := range employees {
		empid, _ := strconv.Atoi(employee.ID)
		if empid == id {
			employees[i] = updatedEmployee // Update the employee details.
			updated = true
			break
		}
	}

	// If the employee with the specified ID was not found, return an error.
	if !updated {
		return fmt.Errorf("Employee with ID %v not found", id)
	}

	// Open the CSV file for writing.
	file, err := os.OpenFile(storage.CsvFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write the updated employee data to the CSV file.
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
