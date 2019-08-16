package error

// UnexpectedError return a new UnexpectedError instance
func UnexpectedError() *Error {
	return &Error{
		code:    999,
		message: "Unexpected Error",
	}
}

// LoginError return a new LoginError instance
func LoginError() *Error {
	return &Error{
		code:    1,
		message: "Login Error, invalid username or password",
	}
}

// ValidateError return a new ValidateError instance
func ValidateError(msg string) *Error {
	return &Error{
		code:    2,
		message: "Validate Error: " + msg,
	}
}
