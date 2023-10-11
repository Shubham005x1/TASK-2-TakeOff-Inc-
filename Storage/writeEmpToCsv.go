package storage

import (
	"TakeOff_Task-2/models"
	"encoding/csv"
	"os"
)

func WriteEmployeeToCSV(employee models.Employee) error {
	file, err := os.OpenFile(CsvFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	row := []string{
		employee.ID,
		employee.FirstName,
		employee.LastName,
		employee.Email,
		employee.Password,
		employee.PhoneNo,
		employee.Role,
	}
	err = writer.Write(row)
	if err != nil {
		return err
	}
	return nil
}
