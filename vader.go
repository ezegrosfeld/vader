package vader

// Generic returns a generic error
func Generic(msg string) Error {
	e := err{msg: msg, code: 0, group: generic, kind: generic}
	return &e
}

// NotFound returns a not found error
func NotFound(msg string) Error {
	e := err{msg: msg, code: notFoundCode, group: generic, kind: notFound}
	return &e
}

// BadRequest returns a bad request error
func BadRequest(msg string) Error {
	e := err{msg: msg, code: badRequestCode, group: generic, kind: badRequest}
	return &e
}

// InternalError returns an internal error
func InternalError(msg string) Error {
	e := err{msg: msg, code: internalCode, group: generic, kind: internal}
	return &e
}

// AlredyExistsError returns an alredy exists error
func AlredyExistsError(msg string) Error {
	e := err{msg: msg, code: alreadyExistsCode, group: generic, kind: alreadyExists}
	return &e
}

// UnauthorizedError returns an unauthorized error
func UnauthorizedError(msg string) Error {
	e := err{msg: msg, code: unauthorizedCode, group: generic, kind: unauthorized}
	return &e
}

// UnprocessableEntityError returns an unprocessable entity error
func UnprocessableEntityError(msg string) Error {
	e := err{msg: msg, code: unprocessableEntityCode, group: generic, kind: unprocessableEntity}
	return &e
}

// NewError returns a new error with status code
func NewError(msg string, status int) Error {
	e := err{msg: msg, code: status, group: generic, kind: generic}
	return &e
}
