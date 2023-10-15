package storage

import (
	"TakeOff_Task-2/models"
	"encoding/csv"
	"os"
)

// WriteEmployeeToCSV writes employee data to the CSV file.
func WriteEmployeeToCSV(employee models.Employee) error {
	// Open the CSV file for writing or create it if it doesn't exist.
	file, err := os.OpenFile(CsvFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	// Initialize a CSV writer.
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Prepare a row of data from the employee struct.
	row := []string{
		employee.ID,
		employee.FirstName,
		employee.LastName,
		employee.Email,
		employee.Password,
		employee.PhoneNo,
		employee.Role,
	}

	// Write the row to the CSV file.
	err = writer.Write(row)
	if err != nil {
		return err
	}

	return nil
}
