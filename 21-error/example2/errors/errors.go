package errors

type CustomError interface {
	Error() string
}

type NewError struct {
	Message string
}

