package error

// Error the basic error type
type Error struct {
	code    int
	message string
}

// Error return error message
func (err *Error) Error() string {
	return err.message
}

// Code return error code
func (err *Error) Code() int {
	return err.code
}
