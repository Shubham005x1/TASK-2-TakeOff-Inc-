package validations

import (
	"regexp"
	"strings"
	"unicode"
)

// IsValidEmail checks if the provided string is a valid email address.
func IsValidEmail(email string) error {
	if strings.Contains(email, "@") && strings.Contains(email, ".") {
		emailRegex := regexp.MustCompile(`^[a-zA-Z0-9.]+@[a-zA-Z0-9]+.[a-zA-Z]{2,}$`)
		if emailRegex.MatchString(email) {
			return nil
		}
	}
	return CustomError{message: "=>Email should contain (@ and .) please enter valid email"}
}

// IsValidSalary checks if the provided salary is a valid positive value.
func IsValidSalary(salary float64) error {

	if salary > 0 {
		return nil
	}
	return CustomError{message: "=>Salary cannot be negative please enter valid salary!"}

}

// IsValidRole checks if the provided role is one of the allowed roles (admin, manager, developer, tester).
func IsValidRole(role string) error {
	allowedRoles := []string{"admin", "manager", "developer", "tester"}
	lowercaseRole := strings.ToLower(role)
	for _, allowedRole := range allowedRoles {
		if lowercaseRole == allowedRole {
			return nil
		}
	}

	return CustomError{message: "=>Invalid Role Please Enter Valid Role"}
}

// IsValidEntry checks if the provided name is a valid entry (not empty and does not contain numbers).
func IsValidEntry(name string) error {
	if name == "" || ValidNameEntry(name) {
		return CustomError{message: "=>Field cannot be empty or cannot have Number in it please provide the valid value "}
	}
	return nil
}

// validNameEntry checks if the provided name contains any non-letter characters.
func ValidNameEntry(name string) bool {
	for _, char := range name {
		if !unicode.IsLetter(char) {
			return true
		}
	}
	return false

}

// IsNumberValid checks if the provided string is a valid numeric value (at least 10 digits and no characters).
func IsNumberValid(num string) error {
	if len(num) < 10 || validNumberEntry(num) {
		return CustomError{message: "=>Number cannot be less than 10 digit and cannot contain character Please try again"}
	}
	return nil
}

// validNumberEntry checks if the provided string contains any non-numeric characters.
func validNumberEntry(name string) bool {
	for _, char := range name {
		if unicode.IsLetter(char) {
			return true
		}
	}
	return false

}
