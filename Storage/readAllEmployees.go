package storage

import (
	"TakeOff_Task-2/models"
	"encoding/csv"
	"os"
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
