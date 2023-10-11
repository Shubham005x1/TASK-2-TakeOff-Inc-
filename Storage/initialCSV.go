package storage

import "os"

var CsvFilePath = "employee_data.csv"

func InitializeCSV() error {

	file, err := os.Create(CsvFilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return nil
}
