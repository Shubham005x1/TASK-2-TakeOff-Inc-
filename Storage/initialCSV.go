package storage

import "os"

// CsvFilePath is the path to the CSV file used for storing employee data.
var CsvFilePath = "employee_data.csv"

// InitializeCSV creates a new CSV file for storing employee data.
func InitializeCSV() error {
	// Create a new CSV file or overwrite it if it already exists.
	file, err := os.Create(CsvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()

	// The CSV file has been successfully created.
	return nil
}
