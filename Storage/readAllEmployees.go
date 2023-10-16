package storage

import (
	"TakeOff_Task-2/models"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

// ReadAllEmployees reads all employee records from the CSV file and returns a slice of Employee structs.
func ReadAllEmployees() ([]models.Employee, error) {
	// Open the CSV file for reading.
	file, err := os.Open(CsvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize a CSV reader.
	reader := csv.NewReader(file)

	// Read all records from the CSV file.
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	var employees []models.Employee

	// Iterate through the records and create Employee structs.
	for _, value := range records {
		emp := models.Employee{
			ID:        value[0],
			FirstName: value[1],
			LastName:  value[2],
			Email:     value[3],
			Password:  value[4],
			PhoneNo:   value[5],
			Role:      value[6],
		}
		employees = append(employees, emp)
	}

	return employees, nil
}

func GetEmployeeByID(id int) (models.Employee, error) {
	employees, err := ReadAllEmployees()
	if err != nil {
		return models.Employee{}, err
	}

	for _, employee := range employees {
		empID, _ := strconv.Atoi(employee.ID)
		if empID == id {
			return employee, nil
		}
	}

	return models.Employee{}, fmt.Errorf("Employee with ID %d not found", id)
}
