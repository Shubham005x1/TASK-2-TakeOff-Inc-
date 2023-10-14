package models

type Employee struct {
	ID        string `json:"id" csv:"id"`
	FirstName string `json:"first_name" csv:"first_name"`
	LastName  string `json:"last_name" csv:"last_name"`
	Email     string `json:"email" csv:"email"`
	Password  string `json:"password" csv:"password"`
	PhoneNo   string `json:"phoneNo" csv:"phoneNo"`
	Role      string `json:"role" csv:"role"`
}
