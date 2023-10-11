package storage

import (
	"TakeOff_Task-2/models"
	"encoding/csv"
	"os"
)

func ReadAllEmployees() ([]models.Employee, error) {
	file, err := os.Open(CsvFilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}
	var employees []models.Employee
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
