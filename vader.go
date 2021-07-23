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
func Internal(msg string) Error {
	e := err{msg: msg, code: internalCode, group: generic, kind: internal}
	return &e
}

// AlredyExistsError returns an alredy exists error
func AlredyExists(msg string) Error {
	e := err{msg: msg, code: alreadyExistsCode, group: generic, kind: alreadyExists}
	return &e
}

// UnauthorizedError returns an unauthorized error
func Unauthorized(msg string) Error {
	e := err{msg: msg, code: unauthorizedCode, group: generic, kind: unauthorized}
	return &e
}

// UnprocessableEntityError returns an unprocessable entity error
func UnprocessableEntity(msg string) Error {
	e := err{msg: msg, code: unprocessableEntityCode, group: generic, kind: unprocessableEntity}
	return &e
}

// Forbidden returns a forbidden error
func Forbidden(msg string) Error {
	e := err{msg: msg, code: forbiddenCode, group: generic, kind: forbidden}
	return &e
}

// Conflict returns a conflict error
func Conflict(msg string) Error {
	e := err{msg: msg, code: conflictCode, group: generic, kind: conflict}
	return &e
}

// NewError returns a new error with status code
func NewError(msg string, status int) Error {
	e := err{msg: msg, code: status, group: generic, kind: generic}
	return &e
}
