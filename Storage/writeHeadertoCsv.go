package storage

import (
	"TakeOff_Task-2/models"
	"encoding/csv"
	"os"
	"reflect"
)

// WriteHeadersToCSV writes the headers of the CSV file based on the struct tags.
func WriteHeadersToCSV() error {
	// Open the CSV file for writing or create it if it doesn't exist.
	file, err := os.OpenFile(CsvFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Get the type of the Employee struct to extract field tags.
	employeeType := reflect.TypeOf(models.Employee{})

	var headers []string

	// Iterate through the fields of the struct and extract CSV tags as headers.
	for i := 0; i < employeeType.NumField(); i++ {
		field := employeeType.Field(i)
		tag := field.Tag.Get("csv")
		headers = append(headers, tag)
	}

	// Write the headers to the CSV file.
	if err := writer.Write(headers); err != nil {
		return err
	}

	return nil
}
