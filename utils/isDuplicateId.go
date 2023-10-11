package utils

import (
	storage "TakeOff_Task-2/Storage"
	"strconv"
)

func IsDuplicateID(id int) bool {
	employees, _ := storage.ReadAllEmployees()

	for _, employee := range employees {
		val, _ := strconv.Atoi(employee.ID)
		if val == id {
			return true
		}
	}

	return false
}
