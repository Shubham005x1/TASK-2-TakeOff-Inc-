package validations

type error interface {
	Error() string
}

type CustomError struct {
	message string
}

func (e CustomError) Error() string {
	return e.message
}
