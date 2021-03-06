package errors

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

// RevokeTokenError return a new RevokeTokenError instance
func RevokeTokenError() *Error {
	return &Error{
		code:    3,
		message: "Revoking token",
	}
}

// NotFoundError return a new NotFoundError instance
func NotFoundError(msg string) *Error {
	return &Error{
		code:    4,
		message: "Not Found: " + msg,
	}
}
