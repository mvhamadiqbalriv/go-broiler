package exception

type DuplicateError struct {
	Error string
}

func NewDuplicateError(error string) DuplicateError {
	return DuplicateError{Error: error}
}