package storage

import (
	"TakeOff_Task-2/models"
	"encoding/csv"
	"os"
	"reflect"
)

func WriteHeadersToCSV() error {
	file, err := os.OpenFile(CsvFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := csv.NewWriter(file)
	defer writer.Flush()

	employeeType := reflect.TypeOf(models.Employee{})

	var headers []string

	for i := 0; i < employeeType.NumField(); i++ {
		field := employeeType.Field(i)
		tag := field.Tag.Get("csv")
		headers = append(headers, tag)
	}

	if err := writer.Write(headers); err != nil {
		return err
	}

	return nil
}
