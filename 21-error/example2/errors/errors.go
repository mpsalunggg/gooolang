package errors

// import "fmt"

type CustomError interface {
	Error() string
}

type NewError struct {
	Message string
}

func (e NewError) Error() string {
	return e.Message
}

func CallError(message string) NewError {
	return NewError{
		Message: message,
	}
}

func ErrorInputBalance() CustomError {
	return CallError("Sorry your input is not valid")
}

func ErrorInvalidMenu() error {
	return CallError("Id menu is not valid!")
}
